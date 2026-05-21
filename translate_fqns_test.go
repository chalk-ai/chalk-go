package chalk

import (
	"testing"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/stretchr/testify/require"
)

func TestTranslateFqnsOnlineQueryResult(t *testing.T) {
	t.Parallel()

	data := []featureResultSerialized{
		{Field: "user.count__86400__", Value: float64(10)},
		{Field: "user.count__3600__", Value: float64(5)},
		{Field: "user.count__all__", Value: float64(100)},
		{Field: "user.id", Value: "u1"},
		{Field: "user.score__86400__@2", Value: float64(0.9)},
	}

	deserialized, err := deserializeFeatureResults(data)
	require.NoError(t, err)
	require.Equal(t, "user.count__86400__", deserialized[0].Field)

	for i := range deserialized {
		deserialized[i].Field = internal.TranslateWindowedFqn(deserialized[i].Field)
	}

	require.Equal(t, "user.count[1d]", deserialized[0].Field)
	require.Equal(t, "user.count[1h]", deserialized[1].Field)
	require.Equal(t, "user.count[all]", deserialized[2].Field)
	require.Equal(t, "user.id", deserialized[3].Field)
	require.Equal(t, "user.score[1d]@2", deserialized[4].Field)
}

func TestTranslateFqnsArrowTable(t *testing.T) {
	t.Parallel()

	alloc := memory.NewGoAllocator()
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "user.count__86400__", Type: arrow.PrimitiveTypes.Int64},
		{Name: "user.count__3600__", Type: arrow.PrimitiveTypes.Int64},
		{Name: "user.id", Type: arrow.BinaryTypes.String},
	}, nil)

	b := array.NewRecordBuilder(alloc, schema)
	defer b.Release()
	b.Field(0).(*array.Int64Builder).AppendValues([]int64{1, 2}, nil)
	b.Field(1).(*array.Int64Builder).AppendValues([]int64{3, 4}, nil)
	b.Field(2).(*array.StringBuilder).AppendValues([]string{"a", "b"}, nil)
	rec := b.NewRecord()
	defer rec.Release()

	tbl := array.NewTableFromRecords(schema, []arrow.Record{rec})
	defer tbl.Release()

	renamed := renameArrowTableColumns(tbl, internal.TranslateWindowedFqn)
	defer renamed.Release()

	require.Equal(t, int64(2), renamed.NumRows())
	require.Equal(t, int64(3), renamed.NumCols())
	require.Equal(t, "user.count[1d]", renamed.Schema().Field(0).Name)
	require.Equal(t, "user.count[1h]", renamed.Schema().Field(1).Name)
	require.Equal(t, "user.id", renamed.Schema().Field(2).Name)
}

func TestWithTranslateFqnsChaining(t *testing.T) {
	t.Parallel()
	p := OnlineQueryParams{}.
		WithInput("user.id", 1).
		WithOutputs("user.count__86400__").
		WithTranslateFqns(true)
	require.True(t, p.underlying.TranslateFqns)
}
