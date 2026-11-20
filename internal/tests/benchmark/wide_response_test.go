package benchmark

import (
	"fmt"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/chalk-ai/chalk-go"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	assert "github.com/stretchr/testify/require"
)

// wideColumns describes the table: kind is s(tring id), f(loat, dense),
// c(ount, ~98.5% null), i(nt), b(ool).
var wideColumns = []struct{ name, kind string }{
	{"wide_row.id", "s"},
	{"wide_row.score_0", "f"},
	{"wide_row.score_1", "f"},
	{"wide_row.score_2", "f"},
	{"wide_row.score_3", "f"},
	{"wide_row.score_4", "f"},
	{"wide_row.score_5", "f"},
	{"wide_row.score_6", "f"},
	{"wide_row.score_7", "f"},
	{"wide_row.score_8", "f"},
	{"wide_row.score_9", "f"},
	{"wide_row.score_10", "f"},
	{"wide_row.score_11", "f"},
	{"wide_row.score_12", "f"},
	{"wide_row.score_13", "f"},
	{"wide_row.score_14", "f"},
	{"wide_row.score_15", "f"},
	{"wide_row.score_16", "f"},
	{"wide_row.score_17", "f"},
	{"wide_row.score_18", "f"},
	{"wide_row.score_19", "f"},
	{"wide_row.score_20", "f"},
	{"wide_row.score_21", "f"},
	{"wide_row.score_22", "f"},
	{"wide_row.score_23", "f"},
	{"wide_row.score_24", "f"},
	{"wide_row.score_25", "f"},
	{"wide_row.score_26", "f"},
	{"wide_row.score_27", "f"},
	{"wide_row.score_28", "f"},
	{"wide_row.score_29", "f"},
	{"wide_row.count_0", "c"},
	{"wide_row.count_1", "c"},
	{"wide_row.count_2", "c"},
	{"wide_row.count_3", "c"},
	{"wide_row.count_4", "c"},
	{"wide_row.count_5", "c"},
	{"wide_row.count_6", "c"},
	{"wide_row.count_7", "c"},
	{"wide_row.count_8", "c"},
	{"wide_row.day_of_week", "i"},
	{"wide_row.offer.rate_0", "f"},
	{"wide_row.offer.rate_1", "f"},
	{"wide_row.offer.rate_2", "f"},
	{"wide_row.offer.rate_3", "f"},
	{"wide_row.offer.rate_4", "f"},
	{"wide_row.offer.rate_5", "f"},
	{"wide_row.offer.rate_6", "f"},
	{"wide_row.offer.rate_7", "f"},
	{"wide_row.offer.rate_8", "f"},
	{"wide_row.offer.rate_9", "f"},
	{"wide_row.offer.rate_10", "f"},
	{"wide_row.offer.rate_11", "f"},
	{"wide_row.offer.rate_12", "f"},
	{"wide_row.offer.rate_13", "f"},
	{"wide_row.offer.rate_14", "f"},
	{"wide_row.offer.rate_15", "f"},
	{"wide_row.offer.rate_16", "f"},
	{"wide_row.offer.rate_17", "f"},
	{"wide_row.offer.rate_18", "f"},
	{"wide_row.offer.rate_19", "f"},
	{"wide_row.offer.restricted", "b"},
}

// wideResponseBytes builds a 1000-row, 60-column Arrow IPC file, as the engine
// returns in OnlineQueryBulkResponse.scalars_data (LZ4-compressed like the
// real thing).
func wideResponseBytes(tb testing.TB, rows int) []byte {
	tb.Helper()
	fields := make([]arrow.Field, len(wideColumns))
	for i, c := range wideColumns {
		var t arrow.DataType
		switch c.kind {
		case "s":
			t = arrow.BinaryTypes.LargeString
		case "f":
			t = arrow.PrimitiveTypes.Float64
		case "b":
			t = arrow.FixedWidthTypes.Boolean
		default:
			t = arrow.PrimitiveTypes.Int64
		}
		fields[i] = arrow.Field{Name: c.name, Type: t, Nullable: true}
	}
	schema := arrow.NewSchema(fields, nil)
	b := array.NewRecordBuilder(memory.DefaultAllocator, schema)
	defer b.Release()
	for r := 0; r < rows; r++ {
		for i, c := range wideColumns {
			switch c.kind {
			case "s":
				b.Field(i).(*array.LargeStringBuilder).Append(fmt.Sprintf("id-%08d", r*7919))
			case "f":
				b.Field(i).(*array.Float64Builder).Append(float64(r*(i+1)) / 7)
			case "b":
				b.Field(i).(*array.BooleanBuilder).Append(r%2 == 0)
			case "i":
				b.Field(i).(*array.Int64Builder).Append(int64(r % 7))
			case "c":
				if r%67 == 0 {
					b.Field(i).(*array.Int64Builder).Append(int64(r + i))
				} else {
					b.Field(i).AppendNull()
				}
			}
		}
	}
	rec := b.NewRecordBatch()
	defer rec.Release()
	bs, err := internal.TableToBytes(array.NewTableFromRecords(schema, []arrow.RecordBatch{rec}), memory.DefaultAllocator)
	assert.NoError(tb, err)
	return bs
}

func wideResultFrom(tb testing.TB, data []byte) *chalk.GRPCOnlineQueryBulkResult {
	res, err := chalk.NewGRPCOnlineQueryBulkResult(&commonv1.OnlineQueryBulkResponse{ScalarsData: data})
	assert.NoError(tb, err)
	return res
}

func wideResult(tb testing.TB, rows int) *chalk.GRPCOnlineQueryBulkResult {
	return wideResultFrom(tb, wideResponseBytes(tb, rows))
}

func checkWide(tb testing.TB, out []fixtures.WideRow, rows int) {
	assert.Len(tb, out, rows)
	last := rows - 1
	assert.Equal(tb, fmt.Sprintf("id-%08d", last*7919), *out[last].Id)
	assert.NotNil(tb, out[last].Offer)
	assert.Equal(tb, float64(last*2)/7, *out[last].Score0)
	assert.Nil(tb, out[1].Count0)
	assert.NotNil(tb, out[0].Count0)
}

// Decoding the Arrow bytes and unmarshalling into structs, the typical typed path.
func BenchmarkWideUnmarshalInto(b *testing.B) {
	data := wideResponseBytes(b, 1000)
	b.ReportAllocs()
	for b.Loop() {
		// A fresh result per iteration: each response is decoded once.
		res := wideResultFrom(b, data)
		var out []fixtures.WideRow
		assert.NoError(b, res.UnmarshalInto(&out))
		if b.N == 1 {
			checkWide(b, out, 1000)
		}
	}
}

func BenchmarkWideUnmarshalIntoParallel(b *testing.B) {
	data := wideResponseBytes(b, 1000)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res := wideResultFrom(b, data)
			var out []fixtures.WideRow
			assert.NoError(b, res.UnmarshalInto(&out))
		}
	})
}

// GetRow is called per row by callers that walk a bulk result (10 rows sampled here).
func BenchmarkWideGetRow(b *testing.B) {
	data := wideResponseBytes(b, 1000)
	b.ReportAllocs()
	for b.Loop() {
		// One response, then every row read via GetRow.
		res := wideResultFrom(b, data)
		for i := 0; i < 1000; i += 100 {
			_, err := res.GetRow(i)
			assert.NoError(b, err)
		}
	}
}

func TestWideUnmarshalCorrect(t *testing.T) {
	res := wideResult(t, 300)
	var out []fixtures.WideRow
	assert.NoError(t, res.UnmarshalInto(&out))
	checkWide(t, out, 300)
	row, err := res.GetRow(299)
	assert.NoError(t, err)
	assert.Equal(t, float64(299*2)/7, row.Features["wide_row.score_0"].Value)
	_, err = res.GetRow(300)
	assert.Error(t, err)
}
