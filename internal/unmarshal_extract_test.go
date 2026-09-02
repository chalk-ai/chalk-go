package internal

import (
	"fmt"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
)

// buildExtractTestTable builds a scalars table shaped like a bulk query result:
// feature columns, the __id__ pkey column, and a result-metadata struct column.
// tsFirst places a skipped "__ts__" column ahead of the features so that the
// feature column indices are not simply 0..n-1.
func buildExtractTestTable(nRows int, tsFirst bool) arrow.Table {
	alloc := memory.NewGoAllocator()
	var fields []arrow.Field
	if tsFirst {
		fields = append(fields, arrow.Field{Name: "__ts__", Type: arrow.PrimitiveTypes.Int64, Nullable: true})
	}
	for i := 0; i < 13; i++ {
		fields = append(fields, arrow.Field{Name: fmt.Sprintf("user.f%02d", i), Type: arrow.PrimitiveTypes.Int64, Nullable: true})
	}
	fields = append(fields, arrow.Field{Name: "__id__", Type: arrow.PrimitiveTypes.Int64, Nullable: true})
	fields = append(fields, arrow.Field{
		Name: metadataPrefix + "user.f00",
		Type: arrow.StructOf(
			arrow.Field{Name: "source_type", Type: arrow.BinaryTypes.String, Nullable: true},
			arrow.Field{Name: "source_id", Type: arrow.BinaryTypes.String, Nullable: true},
			arrow.Field{Name: "resolver_fqn", Type: arrow.BinaryTypes.String, Nullable: true},
		), Nullable: true,
	})

	schema := arrow.NewSchema(fields, nil)
	b := array.NewRecordBuilder(alloc, schema)
	defer b.Release()
	for c := range fields {
		switch fb := b.Field(c).(type) {
		case *array.Int64Builder:
			for r := 0; r < nRows; r++ {
				fb.Append(int64(c*100000 + r))
			}
		case *array.StructBuilder:
			for r := 0; r < nRows; r++ {
				fb.Append(true)
				fb.FieldBuilder(0).(*array.StringBuilder).Append("online_store")
				fb.FieldBuilder(1).(*array.StringBuilder).Append(fmt.Sprintf("src-%d", r))
				fb.FieldBuilder(2).(*array.StringBuilder).Append("user.resolve_f00")
			}
		}
	}
	rec := b.NewRecordBatch()
	defer rec.Release()
	return array.NewTableFromRecords(schema, []arrow.RecordBatch{rec})
}

// Non-feature columns must be skipped without shifting the feature columns,
// even when a skipped column precedes them in the schema.
func TestVerifyColumnMapping(t *testing.T) {
	for _, tsFirst := range []bool{false, true} {
		rows, _, err := ExtractFeaturesFromTable(buildExtractTestTable(1000, tsFirst), false)
		if err != nil {
			t.Fatalf("tsFirst=%v: %v", tsFirst, err)
		}
		if len(rows) != 1000 {
			t.Fatalf("tsFirst=%v: got %d rows", tsFirst, len(rows))
		}
		for _, bad := range []string{"__ts__", "__id__", metadataPrefix + "user.f00"} {
			if _, ok := rows[0][bad]; ok {
				t.Errorf("tsFirst=%v: non-feature column %q leaked into row map", tsFirst, bad)
			}
		}
		if len(rows[0]) != 13 {
			t.Errorf("tsFirst=%v: want 13 features, got %d", tsFirst, len(rows[0]))
		}
		// Spot-check a value is read from the right column, not a shifted one.
		offset := 0
		if tsFirst {
			offset = 1
		}
		for i := 0; i < 13; i++ {
			name := fmt.Sprintf("user.f%02d", i)
			want := int64((i+offset)*100000 + 7)
			if got := rows[7][name]; got != want {
				t.Errorf("tsFirst=%v: rows[7][%s] = %v, want %v", tsFirst, name, got, want)
			}
		}
	}
}

// The primary key must be resolved for every row, not just the first row of
// each chunk, and concurrent chunk workers must not mutate shared state.
func TestVerifyPkeyPerRow(t *testing.T) {
	_, meta, err := ExtractFeaturesFromTable(buildExtractTestTable(1000, false), false)
	if err != nil {
		t.Fatal(err)
	}
	if len(meta) != 1000 {
		t.Fatalf("got %d meta entries", len(meta))
	}
	missing := 0
	for r, m := range meta {
		fm, ok := m["user.f00"]
		if !ok {
			t.Fatalf("row %d: no metadata for user.f00", r)
		}
		want := int64(13*100000 + r) // __id__ is column 13
		if fm.Pkey == nil {
			missing++
		} else if fm.Pkey != want {
			t.Errorf("row %d: Pkey = %v, want %v", r, fm.Pkey, want)
		}
		if fm.ResolverFqn != "user.resolve_f00" {
			t.Errorf("row %d: ResolverFqn = %q", r, fm.ResolverFqn)
		}
	}
	if missing > 0 {
		t.Errorf("%d/1000 rows had a nil Pkey", missing)
	}
}
