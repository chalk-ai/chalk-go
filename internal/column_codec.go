package internal

import (
	"reflect"
	"time"
	"unsafe"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

// disableColumnCodecs forces the row-wise path; for tests that compare the two.
var disableColumnCodecs bool

// columnCodecMinRows is the smallest record (in rows) handled column-wise.
// Below it the row-wise codecs are cheaper because the column path pays a
// fixed cost per column. Tests lower it to exercise the column path on tiny
// tables.
var columnCodecMinRows = 8

// ColumnCodec unmarshals a whole Arrow column into the matching field of
// every row struct in one pass. It is the column-major counterpart of Codec:
// it avoids a closure call, a reflect.Value box and a heap allocation per
// cell, which dominate the cost of unmarshalling wide, many-row responses.
//
// rows is the destination slice of structs, rowOffset is the index in rows of
// the first row of arr, and rootFieldIdx is the index of the namespace struct
// within each row struct (-1 when the row struct itself is the namespace).
type ColumnCodec func(rows reflect.Value, rowOffset int, rootFieldIdx int, arr arrow.Array) error

// codecPair holds the row-wise codec, which handles every supported type, and
// an optional column-wise codec for the common scalar cases. Both are derived
// from the same inputs so they are memoized together.
//
// The memo is keyed by column name only, so one pair can be handed to a
// different struct type than it was built for. The row-wise codec fails safely
// then (reflect panics are recovered into an error), but the column codec
// stores through unsafe pointers using that struct's layout, so it records the
// type it was built for and must only be used with that type.
type codecPair struct {
	row     Codec
	col     ColumnCodec // nil when the column has no fast path
	forType reflect.Type
}

// columnCodecFor returns the column codec if it is valid for structType.
func (p *codecPair) columnCodecFor(structType reflect.Type) ColumnCodec {
	if p.forType != structType {
		return nil
	}
	return p.col
}

// columnValues returns one value per row of arr. owned reports whether the
// slice was freshly allocated (true) or aliases Arrow's memory (false), in
// which case it must be copied before pointers into it are handed out.
type columnValues[T any] func(arr arrow.Array) (vals []T, owned bool)

func numericValues[T any, A interface{ arrow.Array }](values func(A) []T) columnValues[T] {
	return func(arr arrow.Array) ([]T, bool) { return values(arr.(A)), false }
}

func derivedValues[T any](n func(arr arrow.Array) int, get func(arr arrow.Array, i int) T) columnValues[T] {
	return func(arr arrow.Array) ([]T, bool) {
		vals := make([]T, n(arr))
		for i := range vals {
			if arr.IsValid(i) {
				vals[i] = get(arr, i)
			}
		}
		return vals, true
	}
}

func buildColumnCodec[T any](
	fieldType reflect.Type,
	indices []int,
	init InitFeatureFunc,
	values columnValues[T],
) ColumnCodec {
	isPtr := fieldType == reflect.TypeFor[*T]()
	if !isPtr && fieldType != reflect.TypeFor[T]() {
		return nil
	}
	return func(rows reflect.Value, rowOffset int, rootFieldIdx int, arr arrow.Array) error {
		n := arr.Len()
		vals, owned := values(arr)
		if len(vals) < n {
			// Defensive: never index past what Arrow gave us.
			n = len(vals)
		}
		var slab []T
		if isPtr {
			switch {
			case owned:
				slab = vals
			case arr.NullN()*2 < n:
				// Mostly valid: one allocation for the whole column rather than one per cell.
				slab = make([]T, n)
				copy(slab, vals[:n])
			}
		}
		for r := 0; r < n; r++ {
			sv := rows.Index(rowOffset + r)
			if rootFieldIdx >= 0 {
				sv = sv.Field(rootFieldIdx)
			}
			// Like the row-wise codec, has-one structs are initialized even
			// when the value is null.
			if init != nil {
				var err error
				if sv, err = init(sv); err != nil {
					return err
				}
			}
			if arr.IsNull(r) {
				continue
			}
			if isPtr {
				var p *T
				if slab != nil {
					p = &slab[r]
				} else {
					p = new(T)
					*p = vals[r]
				}
				for _, fi := range indices {
					// Field types were verified exactly equal to *T / T and exported
					// when this codec was built, so this is the typed store reflect's
					// Set would do, without its per-call ptrTo lookups.
					*(**T)(unsafe.Pointer(sv.Field(fi).UnsafeAddr())) = p
				}
			} else {
				for _, fi := range indices {
					*(*T)(unsafe.Pointer(sv.Field(fi).UnsafeAddr())) = vals[r]
				}
			}
		}
		return nil
	}
}

func lenOf(arr arrow.Array) int { return arr.Len() }

// newColumnCodec picks a column codec for fieldType and arrowType, or nil when
// there is no fast path and the row-wise codec must be used.
func newColumnCodec(fieldType reflect.Type, arrowType arrow.DataType, indices []int, init InitFeatureFunc) ColumnCodec {
	switch t := arrowType.(type) {
	case *arrow.Int8Type:
		return buildColumnCodec(fieldType, indices, init, numericValues((*array.Int8).Int8Values))
	case *arrow.Int16Type:
		return buildColumnCodec(fieldType, indices, init, numericValues((*array.Int16).Int16Values))
	case *arrow.Int32Type:
		return buildColumnCodec(fieldType, indices, init, numericValues((*array.Int32).Int32Values))
	case *arrow.Int64Type:
		return buildColumnCodec(fieldType, indices, init, numericValues((*array.Int64).Int64Values))
	case *arrow.Uint8Type:
		return buildColumnCodec(fieldType, indices, init, numericValues((*array.Uint8).Uint8Values))
	case *arrow.Uint16Type:
		return buildColumnCodec(fieldType, indices, init, numericValues((*array.Uint16).Uint16Values))
	case *arrow.Uint32Type:
		return buildColumnCodec(fieldType, indices, init, numericValues((*array.Uint32).Uint32Values))
	case *arrow.Uint64Type:
		return buildColumnCodec(fieldType, indices, init, numericValues((*array.Uint64).Uint64Values))
	case *arrow.Float32Type:
		return buildColumnCodec(fieldType, indices, init, numericValues((*array.Float32).Float32Values))
	case *arrow.Float64Type:
		return buildColumnCodec(fieldType, indices, init, numericValues((*array.Float64).Float64Values))
	case *arrow.BooleanType:
		return buildColumnCodec(fieldType, indices, init, derivedValues(lenOf,
			func(a arrow.Array, i int) bool { return a.(*array.Boolean).Value(i) }))
	case *arrow.StringType:
		return buildColumnCodec(fieldType, indices, init, derivedValues(lenOf,
			func(a arrow.Array, i int) string { return a.(*array.String).Value(i) }))
	case *arrow.LargeStringType:
		return buildColumnCodec(fieldType, indices, init, derivedValues(lenOf,
			func(a arrow.Array, i int) string { return a.(*array.LargeString).Value(i) }))
	case *arrow.TimestampType:
		unit := t.TimeUnit()
		return buildColumnCodec(fieldType, indices, init, derivedValues(lenOf,
			func(a arrow.Array, i int) time.Time { return a.(*array.Timestamp).Value(i).ToTime(unit) }))
	case *arrow.Date32Type:
		return buildColumnCodec(fieldType, indices, init, derivedValues(lenOf,
			func(a arrow.Array, i int) time.Time { return a.(*array.Date32).Value(i).ToTime() }))
	case *arrow.Date64Type:
		return buildColumnCodec(fieldType, indices, init, derivedValues(lenOf,
			func(a arrow.Array, i int) time.Time { return a.(*array.Date64).Value(i).ToTime() }))
	}
	return nil
}

// generateColumnCodec mirrors the field resolution done by
// generateUnmarshalValueCodec and returns nil whenever anything is unusual
// (map-typed windowed fields, unknown fields, mismatched types, ...), so that
// the row-wise codec keeps defining behavior for every non-trivial case.
func generateColumnCodec(structType reflect.Type, arrowType arrow.DataType, allMemos *NamespaceMemosT, fqnParts []string) ColumnCodec {
	var init InitFeatureFunc
	if len(fqnParts) > 2 {
		initFunc, leafStructType, err := generateInitFeatureFunc(fqnParts, structType, allMemos)
		if err != nil || initFunc == nil || initFunc == &initFeatureNoOp || leafStructType == nil {
			return nil
		}
		init = *initFunc
		structType = leafStructType
	}
	memo, err := allMemos.LoadOrStore(structType)
	if err != nil {
		return nil
	}
	indices := memo.ResolvedFieldNameToIndices[fqnParts[len(fqnParts)-1]]
	if len(indices) == 0 {
		return nil
	}
	fieldType := structType.Field(indices[0]).Type
	for _, i := range indices {
		// Unexported fields must keep failing the way reflect.Set does.
		if structType.Field(i).Type != fieldType || !structType.Field(i).IsExported() {
			return nil
		}
	}
	if fieldType.Kind() == reflect.Map {
		return nil
	}
	return newColumnCodec(fieldType, arrowType, indices, init)
}

func generateCodecPair(structType reflect.Type, arrowType arrow.DataType, allMemos *NamespaceMemosT, fqn string, fqnParts []string) (*codecPair, error) {
	row, err := generateUnmarshalValueCodec(structType, arrowType, allMemos, fqn, fqnParts)
	if err != nil {
		return nil, err
	}
	return &codecPair{row: *row, col: generateColumnCodec(structType, arrowType, allMemos, fqnParts), forType: structType}, nil
}
