package benchmark

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/chalk-ai/chalk-go"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	assert "github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func getBenchmarkBulkMultiNsPrimitives(b *testing.B, numRows int) func() {
	bulkData := make(map[string]any)
	for i := 0; i < numRows; i++ {
		for j := 1; j <= 40; j++ {
			fqn := fmt.Sprintf("int_features.int_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []float64{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]float64), float64(122.0))

			fqn = fmt.Sprintf("float_features.float_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []float64{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]float64), float64(1.234))

			fqn = fmt.Sprintf("bool_features.bool_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []bool{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]bool), true)

			fqn = fmt.Sprintf("string_features.string_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []string{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]string), fmt.Sprintf("string_val_%d", i))

			fqn = fmt.Sprintf("timestamp_features.timestamp_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []string{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]string), "2024-05-09T22:29:00Z")
		}
	}

	record, err := internal.ColumnMapToRecord(bulkData)
	assert.NoError(b, err)

	table := array.NewTableFromRecords(record.Schema(), []arrow.Record{record})

	res := chalk.OnlineQueryBulkResult{
		ScalarsTable: table,
	}

	assertOnce := sync.Once{}
	benchFunc := func() {
		rootStruct := []struct {
			IntFeatures       fixtures.IntFeatures
			FloatFeatures     fixtures.FloatFeatures
			BoolFeatures      fixtures.BoolFeatures
			StringFeatures    fixtures.StringFeatures
			TimestampFeatures fixtures.TimestampFeatures
		}{}
		assert.NoError(b, res.UnmarshalInto(&rootStruct))
		assertOnce.Do(func() {
			for i := 0; i < numRows; i++ {
				assert.Equal(b, int64(122.0), *rootStruct[i].IntFeatures.Int1)
				assert.Equal(b, int64(122.0), *rootStruct[i].IntFeatures.Int40)
				assert.Equal(b, float64(1.234), *rootStruct[i].FloatFeatures.Float1)
				assert.Equal(b, float64(1.234), *rootStruct[i].FloatFeatures.Float40)
				assert.Equal(b, fmt.Sprintf("string_val_%d", i), *rootStruct[i].StringFeatures.String1)
				assert.Equal(b, fmt.Sprintf("string_val_%d", i), *rootStruct[i].StringFeatures.String40)
				assert.True(b, *rootStruct[i].BoolFeatures.Bool1)
				assert.True(b, *rootStruct[i].BoolFeatures.Bool40)
				assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *rootStruct[i].TimestampFeatures.Timestamp1)
				assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *rootStruct[i].TimestampFeatures.Timestamp40)
			}
		})
	}

	return benchFunc
}

func getBenchmarkMultiNsPrimitives(b *testing.B) func() {
	data := []chalk.FeatureResult{}
	for i := 1; i <= 40; i++ {
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("int_features.int_%d", i),
			Value: float64(122.0),
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("float_features.float_%d", i),
			Value: float64(1.234),
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("bool_features.bool_%d", i),
			Value: true,
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("string_features.string_%d", i),
			Value: "string_val",
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("timestamp_features.timestamp_%d", i),
			Value: "2024-05-09T22:29:00Z",
		})
	}
	res := chalk.OnlineQueryResult{
		Data: data,
	}
	assertOnce := sync.Once{}
	benchFunc := func() {
		rootStruct := struct {
			IntFeatures       fixtures.IntFeatures
			FloatFeatures     fixtures.FloatFeatures
			BoolFeatures      fixtures.BoolFeatures
			StringFeatures    fixtures.StringFeatures
			TimestampFeatures fixtures.TimestampFeatures
		}{}
		assert.NoError(b, res.UnmarshalInto(&rootStruct))

		assertOnce.Do(func() {
			assert.Equal(b, int64(122.0), *rootStruct.IntFeatures.Int1)
			assert.Equal(b, int64(122.0), *rootStruct.IntFeatures.Int40)
			assert.Equal(b, float64(1.234), *rootStruct.FloatFeatures.Float1)
			assert.Equal(b, float64(1.234), *rootStruct.FloatFeatures.Float40)
			assert.Equal(b, "string_val", *rootStruct.StringFeatures.String1)
			assert.Equal(b, "string_val", *rootStruct.StringFeatures.String40)
			assert.True(b, *rootStruct.BoolFeatures.Bool1)
			assert.True(b, *rootStruct.BoolFeatures.Bool40)
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *rootStruct.TimestampFeatures.Timestamp1)
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *rootStruct.TimestampFeatures.Timestamp40)
		})
	}

	return benchFunc
}

func getBenchmarkUnmarshalMultiNsWindowed(t *testing.B) func() {
	newData := []chalk.FeatureResult{}
	windows := []int{60, 300, 3600}
	for i := 1; i <= 13; i++ {
		for _, window := range windows {
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_int_features.int_%d__%d__", i, window),
				Value: float64(122.0),
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_float_features.float_%d__%d__", i, window),
				Value: float64(1.234),
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_bool_features.bool_%d__%d__", i, window),
				Value: true,
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_string_features.string_%d__%d__", i, window),
				Value: "string_val",
			})
			newData = append(newData, chalk.FeatureResult{
				Field: fmt.Sprintf("windowed_timestamp_features.timestamp_%d__%d__", i, window),
				Value: "2024-05-09T22:29:00Z",
			})
		}
	}
	res := chalk.OnlineQueryResult{
		Data: newData,
	}

	assertOnce := sync.Once{}

	benchmarkFunc := func() {
		rootStruct := struct {
			IntFeatures       fixtures.WindowedIntFeatures
			FloatFeatures     fixtures.WindowedFloatFeatures
			BoolFeatures      fixtures.WindowedBoolFeatures
			StringFeatures    fixtures.WindowedStringFeatures
			TimestampFeatures fixtures.WindowedTimestampFeatures
		}{}
		assert.NoError(t, res.UnmarshalInto(&rootStruct))

		assertOnce.Do(func() {
			assert.Equal(t, int64(122.0), *rootStruct.IntFeatures.Int1["1m"])
			assert.Equal(t, int64(122.0), *rootStruct.IntFeatures.Int13["1h"])
			assert.Equal(t, float64(1.234), *rootStruct.FloatFeatures.Float1["1m"])
			assert.Equal(t, float64(1.234), *rootStruct.FloatFeatures.Float13["1h"])
			assert.Equal(t, "string_val", *rootStruct.StringFeatures.String1["1m"])
			assert.Equal(t, "string_val", *rootStruct.StringFeatures.String13["1h"])
			assert.True(t, *rootStruct.BoolFeatures.Bool1["1m"])
			assert.True(t, *rootStruct.BoolFeatures.Bool13["1h"])
			assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *rootStruct.TimestampFeatures.Timestamp1["1m"])
			assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *rootStruct.TimestampFeatures.Timestamp13["1h"])
		})
	}

	return benchmarkFunc
}

func getBenchmarkSingleNs(b *testing.B) func() {
	data := []chalk.FeatureResult{}
	for i := 1; i <= 40; i++ {
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("int_features.int_%d", i),
			Value: float64(122.0),
		})
	}
	res := chalk.OnlineQueryResult{
		Data: data,
	}
	assertOnce := sync.Once{}
	benchFunc := func() {
		intFeatures := fixtures.IntFeatures{}
		assert.NoError(b, res.UnmarshalInto(&intFeatures))

		assertOnce.Do(func() {
			assert.Equal(b, int64(122.0), *intFeatures.Int1)
			assert.Equal(b, int64(122.0), *intFeatures.Int40)
		})
	}

	return benchFunc
}

func getBenchmarkSingleHasOnes(b *testing.B) func() {
	var data []chalk.FeatureResult
	for i := 1; i <= 40; i++ {
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("has_one_root.int_features.int_%d", i),
			Value: float64(122.0),
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("has_one_root.float_features.float_%d", i),
			Value: float64(1.234),
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("has_one_root.bool_features.bool_%d", i),
			Value: true,
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("has_one_root.string_features.string_%d", i),
			Value: "string_val",
		})
		data = append(data, chalk.FeatureResult{
			Field: fmt.Sprintf("has_one_root.timestamp_features.timestamp_%d", i),
			Value: "2024-05-09T22:29:00Z",
		})
	}
	res := chalk.OnlineQueryResult{
		Data: data,
	}
	assertOnce := sync.Once{}
	benchFunc := func() {
		hasOneRoot := fixtures.HasOneRoot{}
		assert.NoError(b, res.UnmarshalInto(&hasOneRoot))

		assertOnce.Do(func() {
			assert.Equal(b, int64(122.0), *hasOneRoot.IntFeatures.Int1)
			assert.Equal(b, int64(122.0), *hasOneRoot.IntFeatures.Int40)
			assert.Equal(b, float64(1.234), *hasOneRoot.FloatFeatures.Float1)
			assert.Equal(b, float64(1.234), *hasOneRoot.FloatFeatures.Float40)
			assert.True(b, *hasOneRoot.BoolFeatures.Bool1)
			assert.True(b, *hasOneRoot.BoolFeatures.Bool40)
			assert.Equal(b, "string_val", *hasOneRoot.StringFeatures.String1)
			assert.Equal(b, "string_val", *hasOneRoot.StringFeatures.String40)
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *hasOneRoot.TimestampFeatures.Timestamp1)
			assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *hasOneRoot.TimestampFeatures.Timestamp40)
		})
	}

	return benchFunc
}

func getBenchmarkBulkSingleNs(b *testing.B) func() {
	bulkData := make(map[string]any)
	for i := 0; i < 100; i++ {
		for j := 1; j <= 40; j++ {
			fqn := fmt.Sprintf("string_features.string_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []string{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]string), fmt.Sprintf("string_val_%d_%d", i, j))
		}
	}

	record, err := internal.ColumnMapToRecord(bulkData)
	assert.NoError(b, err)

	table := array.NewTableFromRecords(record.Schema(), []arrow.Record{record})

	res := chalk.OnlineQueryBulkResult{
		ScalarsTable: table,
	}
	assertOnce := sync.Once{}
	benchFunc := func() {
		stringFeatures := []fixtures.StringFeatures{}
		assert.NoError(b, res.UnmarshalInto(&stringFeatures))

		assertOnce.Do(func() {
			for i := 0; i < 100; i++ {
				assert.Equal(b, fmt.Sprintf("string_val_%d_%d", i, 1), *stringFeatures[i].String1)
				assert.Equal(b, fmt.Sprintf("string_val_%d_%d", i, 40), *stringFeatures[i].String40)
			}
		})
	}

	return benchFunc
}

func getBenchmarkBulkHasOnes(b *testing.B, numRows int) func() {
	bulkData := make(map[string]any)
	for i := 0; i < numRows; i++ {
		for j := 1; j <= 40; j++ {
			fqn := fmt.Sprintf("has_one_root.int_features.int_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []float64{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]float64), float64(122.0))

			fqn = fmt.Sprintf("has_one_root.float_features.float_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []float64{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]float64), float64(1.234))

			fqn = fmt.Sprintf("has_one_root.bool_features.bool_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []bool{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]bool), true)

			fqn = fmt.Sprintf("has_one_root.string_features.string_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []string{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]string), fmt.Sprintf("string_val_%d", i))

			fqn = fmt.Sprintf("has_one_root.timestamp_features.timestamp_%d", j)
			if _, ok := bulkData[fqn]; !ok {
				bulkData[fqn] = []time.Time{}
			}
			bulkData[fqn] = append(bulkData[fqn].([]time.Time), time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC))
		}
	}

	record, err := internal.ColumnMapToRecord(bulkData)
	assert.NoError(b, err)

	table := array.NewTableFromRecords(record.Schema(), []arrow.Record{record})

	res := chalk.OnlineQueryBulkResult{ScalarsTable: table}

	assertOnce := sync.Once{}
	benchFunc := func() {
		roots := []fixtures.HasOneRoot{}
		assert.NoError(b, res.UnmarshalInto(&roots))
		assertOnce.Do(func() {
			for i := 0; i < numRows; i++ {
				assert.Equal(b, int64(122.0), *roots[i].IntFeatures.Int1)
				assert.Equal(b, int64(122.0), *roots[i].IntFeatures.Int40)
				assert.Equal(b, float64(1.234), *roots[i].FloatFeatures.Float1)
				assert.Equal(b, float64(1.234), *roots[i].FloatFeatures.Float40)
				assert.Equal(b, fmt.Sprintf("string_val_%d", i), *roots[i].StringFeatures.String1)
				assert.Equal(b, fmt.Sprintf("string_val_%d", i), *roots[i].StringFeatures.String40)
				assert.True(b, *roots[i].BoolFeatures.Bool1)
				assert.True(b, *roots[i].BoolFeatures.Bool40)
				assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *roots[i].TimestampFeatures.Timestamp1)
				assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *roots[i].TimestampFeatures.Timestamp40)
			}
		})
	}

	return benchFunc
}

func getBenchmarkUnmarshalBulkAllTypes(b *testing.B) func() {
	bulkData := make(map[string]any)

	numRows := 100

	bulkData["all_types.int"] = make([]int, numRows)
	bulkData["all_types.float"] = make([]float64, numRows)
	bulkData["all_types.string"] = make([]string, numRows)
	bulkData["all_types.bool"] = make([]bool, numRows)
	bulkData["all_types.timestamp"] = make([]time.Time, numRows)
	bulkData["all_types.int_list"] = make([][]int, numRows)
	bulkData["all_types.nested_int_pointer_list"] = make([][][]int, numRows)
	bulkData["all_types.nested_int_list"] = make([][][]int, numRows)
	bulkData["all_types.windowed_int__60__"] = make([]int, numRows)
	bulkData["all_types.windowed_int__300__"] = make([]int, numRows)
	bulkData["all_types.windowed_int__3600__"] = make([]int, numRows)
	bulkData["all_types.windowed_list__60__"] = make([][]int, numRows)
	bulkData["all_types.dataclass"] = make([]fixtures.LatLng, numRows)
	bulkData["all_types.dataclass_list"] = make([][]fixtures.LatLng, numRows)
	bulkData["all_types.dataclass_with_list"] = make([]fixtures.FavoriteThings, numRows)
	bulkData["all_types.dataclass_with_nils"] = make([]fixtures.Possessions, numRows)
	bulkData["all_types.dataclass_with_dataclass"] = make([]fixtures.Child, numRows)
	bulkData["all_types.dataclass_with_overrides"] = make([]fixtures.DclassWithOverrides, numRows)
	bulkData["all_types.nested"] = make([]fixtures.LevelOneNest, numRows)

	for i := 0; i < numRows; i++ {
		bulkData["all_types.int"].([]int)[i] = 1
		bulkData["all_types.float"].([]float64)[i] = 1.234
		bulkData["all_types.string"].([]string)[i] = "string_val"
		bulkData["all_types.bool"].([]bool)[i] = true
		bulkData["all_types.timestamp"].([]time.Time)[i] = time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC)
		bulkData["all_types.int_list"].([][]int)[i] = []int{1}
		bulkData["all_types.nested_int_pointer_list"].([][][]int)[i] = [][]int{[]int{1}}
		bulkData["all_types.nested_int_list"].([][][]int)[i] = [][]int{[]int{1}}
		bulkData["all_types.windowed_int__60__"].([]int)[i] = 1
		bulkData["all_types.windowed_int__300__"].([]int)[i] = 2
		bulkData["all_types.windowed_int__3600__"].([]int)[i] = 3
		bulkData["all_types.windowed_list__60__"].([][]int)[i] = []int{4}
		bulkData["all_types.dataclass"].([]fixtures.LatLng)[i] = fixtures.LatLng{Lat: ptr.Ptr(1.0), Lng: ptr.Ptr(1.0)}
		bulkData["all_types.dataclass_list"].([][]fixtures.LatLng)[i] = []fixtures.LatLng{fixtures.LatLng{Lat: ptr.Ptr(1.0), Lng: ptr.Ptr(1.0)}}
		bulkData["all_types.dataclass_with_list"].([]fixtures.FavoriteThings)[i] = fixtures.FavoriteThings{Numbers: &[]int64{1}}
		bulkData["all_types.dataclass_with_dataclass"].([]fixtures.Child)[i] = fixtures.Child{Name: ptr.Ptr("child"), Mom: &fixtures.Parent{Name: ptr.Ptr("mom-1")}, Dad: &fixtures.Parent{Name: ptr.Ptr("dad-1"), Mom: &fixtures.Grandparent{Name: ptr.Ptr("dad-1-mom")}}}
		bulkData["all_types.nested"].([]fixtures.LevelOneNest)[i] = fixtures.LevelOneNest{Id: ptr.Ptr("level-1-id"), Nested: &fixtures.LevelTwoNest{Id: ptr.Ptr("level-2-id")}}
	}

	record, err := internal.ColumnMapToRecord(bulkData)
	assert.NoError(b, err)

	table := array.NewTableFromRecords(record.Schema(), []arrow.Record{record})
	res := chalk.OnlineQueryBulkResult{ScalarsTable: table}

	assertOnce := sync.Once{}
	return func() {
		allTypes := []fixtures.AllTypes{}
		assert.NoError(b, res.UnmarshalInto(&allTypes))
		assertOnce.Do(func() {
			assert.Equal(b, int64(numRows), table.NumRows())
			numSamples := 10
			interval := numRows / numSamples
			for i := 0; i < numRows; i = i + interval {
				assert.Equal(b, int64(1), *allTypes[i].Int)
				assert.Equal(b, float64(1.234), *allTypes[i].Float)
				assert.Equal(b, "string_val", *allTypes[i].String)
				assert.True(b, *allTypes[i].Bool)
				assert.Equal(b, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *allTypes[i].Timestamp)
				assert.Equal(b, []int64{1}, *allTypes[i].IntList)
				assert.Equal(b, []*[]int64{&[]int64{1}}, *allTypes[i].NestedIntPointerList)
				assert.Equal(b, [][]int64{[]int64{1}}, *allTypes[i].NestedIntList)
				assert.Equal(b, int64(1), *allTypes[i].WindowedInt["1m"])
				assert.Equal(b, int64(2), *allTypes[i].WindowedInt["5m"])
				assert.Equal(b, int64(3), *allTypes[i].WindowedInt["1h"])
				assert.Equal(b, []int64{4}, *allTypes[i].WindowedList["1m"])
				assert.Equal(b, fixtures.LatLng{Lat: ptr.Ptr(1.0), Lng: ptr.Ptr(1.0)}, *allTypes[i].Dataclass)
				assert.Equal(b, []fixtures.LatLng{fixtures.LatLng{Lat: ptr.Ptr(1.0), Lng: ptr.Ptr(1.0)}}, *allTypes[i].DataclassList)
				assert.Equal(b, fixtures.FavoriteThings{Numbers: &[]int64{1}}, *allTypes[i].DataclassWithList)
				assert.Equal(b, fixtures.Child{Name: ptr.Ptr("child"), Mom: &fixtures.Parent{Name: ptr.Ptr("mom-1")}, Dad: &fixtures.Parent{Name: ptr.Ptr("dad-1"), Mom: &fixtures.Grandparent{Name: ptr.Ptr("dad-1-mom")}}}, *allTypes[i].DataclassWithDataclass)
				assert.Equal(b, fixtures.LevelOneNest{Id: ptr.Ptr("level-1-id"), Nested: &fixtures.LevelTwoNest{Id: ptr.Ptr("level-2-id")}}, *allTypes[i].Nested)
			}
		})
	}
}

/*
 * Query: Single
 * Namespaces: Single
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalSingleNsPrimitivesSingle(b *testing.B) {
	benchmark(b, getBenchmarkSingleNs(b))
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Windowed
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalMultiNsWindowedSingle(t *testing.B) {
	benchmark(t, getBenchmarkUnmarshalMultiNsWindowed(t))
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Windowed
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalMultiNsWindowedParallel(t *testing.B) {
	benchmarkParallel(t, getBenchmarkUnmarshalMultiNsWindowed(t))
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalMultiNsPrimitivesSingle(b *testing.B) {
	benchmark(b, getBenchmarkMultiNsPrimitives(b))
}

/*
 * Query: Single
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalMultiNsPrimitivesParallel(b *testing.B) {
	benchmarkParallel(b, getBenchmarkMultiNsPrimitives(b))
}

/*
 * Query: Bulk
 * Namespaces: Single
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalBulkSingleNsPrimitivesSingle(b *testing.B) {
	benchmark(b, getBenchmarkBulkSingleNs(b))
}

/*
 * Query: Bulk
 * Namespaces: Single
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalBulkSingleNsPrimitivesParallel(b *testing.B) {
	benchmarkParallel(b, getBenchmarkBulkSingleNs(b))
}

/*
 * Query: Bulk
 * Namespaces: Single
 * Feature Type: All Types
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalBulkSingleNsAllTypesSingle(b *testing.B) {
	benchmark(b, getBenchmarkUnmarshalBulkAllTypes(b))
}

/*
 * Query: Bulk
 * Namespaces: Single
 * Feature Type: All Types
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalBulkSingleNsAllTypesParallel(b *testing.B) {
	benchmarkParallel(b, getBenchmarkUnmarshalBulkAllTypes(b))
}

/*
 * Query: Bulk
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalBulkMultiNsPrimitivesSingle(b *testing.B) {
	benchmark(b, getBenchmarkBulkMultiNsPrimitives(b, 100))
}

/*
 * Query: Bulk
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalBulkMultiNsPrimitivesParallel(b *testing.B) {
	benchmarkParallel(b, getBenchmarkBulkMultiNsPrimitives(b, 100))
}

/*
 * Query: Bulk (single row)
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Single
 */
func BenchmarkUnmarshalBulkLoneMultiNsPrimitivesSingle(b *testing.B) {
	benchmark(b, getBenchmarkBulkMultiNsPrimitives(b, 1))
}

/*
 * Query: Bulk (single row)
 * Namespaces: Multi
 * Feature Type: Primitives
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalBulkLoneMultiNsPrimitivesParallel(b *testing.B) {
	benchmarkParallel(b, getBenchmarkBulkMultiNsPrimitives(b, 1))
}

/*
 * Query: Single
 * Namespaces: Single
 * Feature Type: Has Ones
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalHasOnes(b *testing.B) {
	benchmarkParallel(b, getBenchmarkSingleHasOnes(b))
}

/*
 * Query: Bulk (single row)
 * Namespaces: Single
 * Feature Type: Has Ones
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalBulkLoneHasOnes(b *testing.B) {
	benchmarkParallel(b, getBenchmarkBulkHasOnes(b, 1))
}

/*
 * Query: Bulk
 * Namespaces: Single
 * Feature Type: Has Ones
 * Protocol: REST
 * Run Type: Parallel
 */
func BenchmarkUnmarshalBulkHasOnes(b *testing.B) {
	benchmarkParallel(b, getBenchmarkBulkHasOnes(b, 100))
}
