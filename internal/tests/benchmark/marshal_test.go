package benchmark

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func getSingleRowPrimitives(b *testing.B) func() {
	colMap := map[string]any{}

	for i := 0; i < 10; i++ {
		colMap[fmt.Sprintf("all_types.int_%d", i)] = []int64{int64(i)}
		colMap[fmt.Sprintf("all_types.float_%d", i)] = []float64{float64(i)}
		colMap[fmt.Sprintf("all_types.string_%d", i)] = []string{fmt.Sprintf("string_%d", i)}
		colMap[fmt.Sprintf("all_types.bool_%d", i)] = []bool{i%2 == 0}
		colMap[fmt.Sprintf("all_types.timestamp_%d", i)] = []time.Time{time.Now()}
	}

	return func() {
		_, err := internal.ColumnMapToRecord(colMap)
		assert.NoError(b, err)
	}
}

func getAllTypes(b *testing.B, numRows int) func() {
	colMap := map[string]any{}

	type customHasMany struct {
		Id *string

		// The following should be kept in parity with the enumeration of
		// fields with all types in the `AllTypes` struct.
		Int    *int64
		Float  *float64
		String *string
	}

	numCols := 3
	for j := 0; j < numCols; j++ {
		colMap[fmt.Sprintf("all_types.int_%d", j)] = []int64{}
		colMap[fmt.Sprintf("all_types.float_%d", j)] = []float64{}
		colMap[fmt.Sprintf("all_types.string_%d", j)] = []string{}
		colMap[fmt.Sprintf("all_types.bool_%d", j)] = []bool{}
		colMap[fmt.Sprintf("all_types.timestamp_%d", j)] = []time.Time{}
		colMap[fmt.Sprintf("all_types.int_list_%d", j)] = [][]int64{}
		colMap[fmt.Sprintf("all_types.nested_int_pointer_list_%d", j)] = [][]*[]int64{}
		colMap[fmt.Sprintf("all_types.nested_int_list_%d", j)] = [][]int64{}
		colMap[fmt.Sprintf("all_types.dataclass_%d", j)] = []*fixtures.LatLng{}
		colMap[fmt.Sprintf("all_types.dataclass_list_%d", j)] = [][]fixtures.LatLng{}
		colMap[fmt.Sprintf("all_types.dataclass_with_list_%d", j)] = []*fixtures.FavoriteThings{}
		colMap[fmt.Sprintf("all_types.dataclass_with_nils_%d", j)] = []*fixtures.Possessions{}
		colMap[fmt.Sprintf("all_types.dataclass_with_dataclass_%d", j)] = []*fixtures.Child{}
		colMap[fmt.Sprintf("all_types.nested_%d", j)] = []*fixtures.LevelOneNest{}
		colMap[fmt.Sprintf("all_types.has_many_%d", j)] = [][]customHasMany{}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < 3; j++ {
			fqn := fmt.Sprintf("all_types.int_%d", j)
			colMap[fqn] = append(colMap[fqn].([]int64), int64(i))
			fqn = fmt.Sprintf("all_types.float_%d", j)
			colMap[fqn] = append(colMap[fqn].([]float64), float64(i))
			fqn = fmt.Sprintf("all_types.string_%d", j)
			colMap[fqn] = append(colMap[fqn].([]string), fmt.Sprintf("string_%d", i))
			fqn = fmt.Sprintf("all_types.bool_%d", j)
			colMap[fqn] = append(colMap[fqn].([]bool), i%2 == 0)
			fqn = fmt.Sprintf("all_types.timestamp_%d", j)
			colMap[fqn] = append(colMap[fqn].([]time.Time), time.Now())
			fqn = fmt.Sprintf("all_types.int_list_%d", j)
			colMap[fqn] = append(colMap[fqn].([][]int64), []int64{int64(i)})
			fqn = fmt.Sprintf("all_types.nested_int_pointer_list_%d", j)
			colMap[fqn] = append(colMap[fqn].([][]*[]int64), []*[]int64{&[]int64{int64(i)}})
			fqn = fmt.Sprintf("all_types.nested_int_list_%d", j)
			colMap[fqn] = append(colMap[fqn].([][]int64), []int64{int64(i)})
			fqn = fmt.Sprintf("all_types.dataclass_%d", j)
			colMap[fqn] = append(colMap[fqn].([]*fixtures.LatLng), &fixtures.LatLng{
				Lat: ptr.Ptr(float64(i)),
				Lng: ptr.Ptr(float64(i)),
			})
			fqn = fmt.Sprintf("all_types.dataclass_list_%d", j)
			colMap[fqn] = append(colMap[fqn].([][]fixtures.LatLng), []fixtures.LatLng{
				{Lat: ptr.Ptr(float64(i)), Lng: ptr.Ptr(float64(i))},
				{Lat: ptr.Ptr(float64(i)), Lng: ptr.Ptr(float64(i))},
			})
			fqn = fmt.Sprintf("all_types.dataclass_with_list_%d", j)
			colMap[fqn] = append(colMap[fqn].([]*fixtures.FavoriteThings), &fixtures.FavoriteThings{
				Numbers: &[]int64{int64(i)},
				Words:   &[]string{fmt.Sprintf("string_%d", i)},
			})
			fqn = fmt.Sprintf("all_types.dataclass_with_nils_%d", j)
			colMap[fqn] = append(colMap[fqn].([]*fixtures.Possessions), &fixtures.Possessions{
				Car:   ptr.Ptr(fmt.Sprintf("car_%d", i)),
				Yacht: ptr.Ptr(fmt.Sprintf("yacht_%d", i)),
			})
			fqn = fmt.Sprintf("all_types.dataclass_with_dataclass_%d", j)
			colMap[fqn] = append(colMap[fqn].([]*fixtures.Child), &fixtures.Child{
				Name: ptr.Ptr(fmt.Sprintf("child_%d", i)),
				Mom: &fixtures.Parent{
					Name: ptr.Ptr(fmt.Sprintf("mom_%d", i)),
				},
			})
			fqn = fmt.Sprintf("all_types.nested_%d", j)
			colMap[fqn] = append(colMap[fqn].([]*fixtures.LevelOneNest), &fixtures.LevelOneNest{
				Id: ptr.Ptr(fmt.Sprintf("id_%d", i)),
				Nested: &fixtures.LevelTwoNest{
					Id: ptr.Ptr(fmt.Sprintf("id_%d", i)),
				},
			})
			fqn = fmt.Sprintf("all_types.has_many_%d", j)
			colMap[fqn] = append(colMap[fqn].([][]customHasMany), []customHasMany{
				{Int: ptr.Ptr(int64(i)), Float: ptr.Ptr(float64(i)), String: ptr.Ptr(fmt.Sprintf("string_%d", i))},
				{Int: ptr.Ptr(int64(i)), Float: ptr.Ptr(float64(i)), String: ptr.Ptr(fmt.Sprintf("string_%d", i))},
				{Int: ptr.Ptr(int64(i)), Float: ptr.Ptr(float64(i)), String: ptr.Ptr(fmt.Sprintf("string_%d", i))},
			})
		}
	}

	return func() {
		_, err := internal.ColumnMapToRecord(colMap)
		assert.NoError(b, err)
	}
}

/*
 * Feature Type: Primitives
 * Rows: Single
 */
func BenchmarkMakeRecordSingleRowPrimitives(b *testing.B) {
	benchmarkParallel(b, getSingleRowPrimitives(b))
}

/*
 * Feature Type: All Types
 * Rows: Single
 */
func BenchmarkMakeRecordSingleRowAllTypes(b *testing.B) {
	benchmarkParallel(b, getAllTypes(b, 1))
}

/*
 * Feature Type: All Types
 * Rows: Many
 */
func BenchmarkMakeRecordManyRowsAllTypes(b *testing.B) {
	benchmarkParallel(b, getAllTypes(b, 20))
}
