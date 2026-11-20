package internal

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/stretchr/testify/require"
)

type ccChild struct {
	Score *float64 `json:"score"`
	Label *string  `json:"label"`
}

type ccUser struct {
	I8      *int8              `json:"i8"`
	I16     *int16             `json:"i16"`
	I32     *int32             `json:"i32"`
	I64     *int64             `json:"i64"`
	Plain   int64              `json:"plain"` // non-pointer field
	U8      *uint8             `json:"u8"`
	U16     *uint16            `json:"u16"`
	U32     *uint32            `json:"u32"`
	U64     *uint64            `json:"u64"`
	F32     *float32           `json:"f32"`
	F64     *float64           `json:"f64"`
	B       *bool              `json:"b"`
	S       *string            `json:"s"`
	LS      *string            `json:"ls"`
	Ts      *time.Time         `json:"ts_us"`
	D32     *time.Time         `json:"d32"`
	Win     map[string]*int64  `json:"win" windows:"1m,5m"`
	Child   *ccChild           `json:"child"`
	Ver     *int64             `json:"ver" versioned:"default(2)"`
	Missing *int64             `json:"missing_in_table"`
	Skipped *float64           `json:"skipped"` // column exists in table but has a different Go type below
	Others  map[string]float64 `json:"others" windows:"1m"`
}

func ccBuildTable(nRows int, records int) arrow.Table {
	alloc := memory.NewGoAllocator()
	f := func(name string, t arrow.DataType) arrow.Field {
		return arrow.Field{Name: name, Type: t, Nullable: true}
	}
	ns := "cc_user."
	fields := []arrow.Field{
		f(ns+"i8", arrow.PrimitiveTypes.Int8), f(ns+"i16", arrow.PrimitiveTypes.Int16),
		f(ns+"i32", arrow.PrimitiveTypes.Int32), f(ns+"i64", arrow.PrimitiveTypes.Int64),
		f(ns+"plain", arrow.PrimitiveTypes.Int64),
		f(ns+"u8", arrow.PrimitiveTypes.Uint8), f(ns+"u16", arrow.PrimitiveTypes.Uint16),
		f(ns+"u32", arrow.PrimitiveTypes.Uint32), f(ns+"u64", arrow.PrimitiveTypes.Uint64),
		f(ns+"f32", arrow.PrimitiveTypes.Float32), f(ns+"f64", arrow.PrimitiveTypes.Float64),
		f(ns+"b", arrow.FixedWidthTypes.Boolean),
		f(ns+"s", arrow.BinaryTypes.String), f(ns+"ls", arrow.BinaryTypes.LargeString),
		f(ns+"ts_us", &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"}),
		f(ns+"d32", arrow.FixedWidthTypes.Date32),
		f(ns+"win__60__", arrow.PrimitiveTypes.Int64), f(ns+"win__300__", arrow.PrimitiveTypes.Int64),
		f(ns+"child.score", arrow.PrimitiveTypes.Float64), f(ns+"child.label", arrow.BinaryTypes.String),
		f(ns+"ver@2", arrow.PrimitiveTypes.Int64),
		f(ns+"unknown_to_struct", arrow.PrimitiveTypes.Int64),
		f(ns+"skipped", arrow.PrimitiveTypes.Float64),
		f(ns+"others__60__", arrow.PrimitiveTypes.Float64),
		f("__ts__", arrow.PrimitiveTypes.Int64),
	}
	schema := arrow.NewSchema(fields, nil)
	var recs []arrow.RecordBatch
	per := nRows / records
	for rec := 0; rec < records; rec++ {
		b := array.NewRecordBuilder(alloc, schema)
		for r := 0; r < per; r++ {
			g := rec*per + r
			for c := range fields {
				// Different null cadence per column, including all-null and no-null.
				null := (c%4 == 0 && g%3 == 0) || (c%7 == 1 && g%10 != 0) || c == 22 || ((c == 18 || c == 19) && g%5 == 0) // child all-null on some rows
				if c == 4 {
					null = false
				}
				if null {
					b.Field(c).AppendNull()
					continue
				}
				switch fb := b.Field(c).(type) {
				case *array.Int8Builder:
					fb.Append(int8(g))
				case *array.Int16Builder:
					fb.Append(int16(g * 3))
				case *array.Int32Builder:
					fb.Append(int32(g * 7))
				case *array.Int64Builder:
					fb.Append(int64(g * 11))
				case *array.Uint8Builder:
					fb.Append(uint8(g))
				case *array.Uint16Builder:
					fb.Append(uint16(g))
				case *array.Uint32Builder:
					fb.Append(uint32(g))
				case *array.Uint64Builder:
					fb.Append(uint64(g))
				case *array.Float32Builder:
					fb.Append(float32(g) / 4)
				case *array.Float64Builder:
					fb.Append(float64(g) / 8)
				case *array.BooleanBuilder:
					fb.Append(g%2 == 0)
				case *array.StringBuilder:
					fb.Append(fmt.Sprintf("s-%d", g))
				case *array.LargeStringBuilder:
					fb.Append(fmt.Sprintf("ls-%d", g))
				case *array.TimestampBuilder:
					fb.Append(arrow.Timestamp(int64(g) * 1_000_000))
				case *array.Date32Builder:
					fb.Append(arrow.Date32(g))
				}
			}
		}
		recs = append(recs, b.NewRecordBatch())
		b.Release()
	}
	return array.NewTableFromRecords(schema, recs)
}

func ccUnmarshal(t *testing.T, table arrow.Table, disable bool) ([]ccUser, error) {
	disableColumnCodecs = disable
	columnCodecMinRows = 1
	defer func() { disableColumnCodecs = false; columnCodecMinRows = 8 }()
	var out []ccUser
	err := UnmarshalTableInto(table, &out)
	return out, err
}

// The column-wise fast path must produce exactly what the row-wise codecs do,
// across types, nulls, has-ones, windows, versioned fields and record batches.
func TestColumnCodecsMatchRowCodecs(t *testing.T) {
	old := TableReaderChunkSize
	TableReaderChunkSize = 400
	defer func() { TableReaderChunkSize = old }()
	for _, tc := range []struct{ rows, records int }{{0, 1}, {1, 1}, {1200, 1}, {1200, 3}, {999, 1}} {
		t.Run(fmt.Sprintf("%d_rows_%d_records", tc.rows, tc.records), func(t *testing.T) {
			table := ccBuildTable(tc.rows/tc.records*tc.records, tc.records)
			defer table.Release()
			want, err := ccUnmarshal(t, table, true)
			require.NoError(t, err)
			got, err := ccUnmarshal(t, table, false)
			require.NoError(t, err)
			require.Equal(t, want, got)
			if tc.rows >= 400 {
				// sanity: the data is really populated, not trivially equal
				require.NotNil(t, got[1].F64)
				require.NotNil(t, got[1].Child)
				require.NotNil(t, got[1].Ts)
			}
		})
	}
}

// A column whose Go field type differs from the Arrow type must fail the
// same way on both paths, and an unexported field must never be written.
func TestColumnCodecsErrorParity(t *testing.T) {
	type mismatch struct {
		A *string `json:"a_mismatch"`
	}
	type unexported struct {
		a *int64
	}
	alloc := memory.NewGoAllocator()
	mk := func(name string) arrow.Table {
		schema := arrow.NewSchema([]arrow.Field{{Name: name, Type: arrow.PrimitiveTypes.Int64, Nullable: true}}, nil)
		b := array.NewRecordBuilder(alloc, schema)
		defer b.Release()
		b.Field(0).(*array.Int64Builder).AppendValues([]int64{1, 2}, nil)
		rec := b.NewRecordBatch()
		return array.NewTableFromRecords(schema, []arrow.RecordBatch{rec})
	}
	columnCodecMinRows = 1
	defer func() { columnCodecMinRows = 8 }()
	for _, disable := range []bool{true, false} {
		disableColumnCodecs = disable
		var m []mismatch
		require.Error(t, UnmarshalTableInto(mk("mismatch.a_mismatch"), &m), "disable=%v", disable)
		var u []unexported
		require.Error(t, UnmarshalTableInto(mk("unexported.a"), &u), "disable=%v", disable)
		for _, row := range u {
			require.Nil(t, row.a, "unexported field must never be written, disable=%v", disable)
		}
	}
	disableColumnCodecs = false
}

func TestExtractFeatureRowMatchesExtractFeatures(t *testing.T) {
	old := TableReaderChunkSize
	TableReaderChunkSize = 7
	defer func() { TableReaderChunkSize = old }()
	for _, tsFirst := range []bool{false, true} {
		table := buildExtractTestTable(40, tsFirst)
		rows, meta, err := ExtractFeaturesFromTable(table, false)
		require.NoError(t, err)
		for i := range rows {
			row, rowMeta, err := ExtractFeatureRowFromTable(table, i, false)
			require.NoError(t, err)
			require.Equal(t, rows[i], row)
			require.Equal(t, meta[i], rowMeta)
		}
		for _, bad := range []int{-1, len(rows), len(rows) + 5} {
			_, _, err := ExtractFeatureRowFromTable(table, bad, false)
			require.ErrorContains(t, err, "out of bounds")
		}
	}
}

var _ = reflect.TypeOf

// The codec memo is keyed by column name only, so two struct types can share
// a memoized codec. The column codec writes through unsafe pointers using the
// layout of the struct it was built for, so it must never be applied to a
// different struct type: the result must be the row path's safe error.
func TestColumnCodecNotReusedAcrossStructTypes(t *testing.T) {
	type structA struct {
		Age *int64 `json:"age"`
	}
	type structB struct {
		Pad [4]int64
		Age *string `json:"age"`
	}
	alloc := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{{Name: "cc_shared_ns.age", Type: arrow.PrimitiveTypes.Int64, Nullable: true}}, nil)
	b := array.NewRecordBuilder(alloc, schema)
	defer b.Release()
	for i := 0; i < 20; i++ {
		b.Field(0).(*array.Int64Builder).Append(int64(i))
	}
	table := array.NewTableFromRecords(schema, []arrow.RecordBatch{b.NewRecordBatch()})

	var a []structA
	require.NoError(t, UnmarshalTableInto(table, &a))
	require.Equal(t, int64(7), *a[7].Age)

	var bs []structB
	err := UnmarshalTableInto(table, &bs)
	require.Error(t, err)
	for _, r := range bs {
		require.Equal(t, [4]int64{}, r.Pad, "nothing may be written into the wrong field")
	}
}
