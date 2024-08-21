package chalk

import (
	"encoding/base64"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	assert "github.com/stretchr/testify/require"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

type unmarshalLatLng struct {
	Lat *float64 `dataclass_field:"true"`
	Lng *float64 `dataclass_field:"true"`
}

type unmarshalUser struct {
	Id *string

	Int *int64

	// Versioned features
	Grade   *int `versioned:"default(2)"`
	GradeV1 *int `versioned:"true"`
	GradeV2 *int `versioned:"true"`

	// Windowed features
	AvgSpend map[string]*float64 `windows:"1m,5m,1h"`

	// Dataclass features
	LatLng *unmarshalLatLng
}

type user struct {
	// TODO: This has to be "user" and not any other name,
	//       otherwise unmarshalling returns a validation error.
	//       But we really shouldn't need this struct to be
	//       named "user" in order to unmarshal features in the
	//       namespace "user".
	Id              *int64
	FavoriteNumbers *[]int64
	FavoriteColors  *[]string
}

func TestOnlineQueryUnmarshalNonBulkAllTypes(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)

	// Mimic JSON deser which returns all numbers as `float64`
	data := []FeatureResult{
		{
			Field: "all_types.int",
			Value: float64(123),
		},
		{
			Field: "all_types.float",
			Value: float64(123),
		},
		{
			Field: "all_types.string",
			Value: "abc",
		},
		{
			Field: "all_types.bool",
			Value: true,
		},
		{
			Field: "all_types.timestamp",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "all_types.int_list",
			Value: []any{float64(1), float64(2), float64(3)},
		},
		{
			Field: "all_types.nested_int_pointer_list",
			Value: []any{[]any{float64(1), float64(2)}, []any{float64(3), float64(4)}},
		},
		{
			Field: "all_types.nested_int_list",
			Value: []any{[]any{float64(1), float64(2)}, []any{float64(3), float64(4)}},
		},
		{
			Field: "all_types.windowed_int__60__",
			Value: 1,
		},
		{
			Field: "all_types.windowed_int__300__",
			Value: 2,
		},
		{
			Field: "all_types.windowed_int__3600__",
			Value: 3,
		},
		{
			Field: "all_types.dataclass",
			Value: []any{float64(1.0), float64(2.0)},
		},
		{
			Field: "all_types.dataclass_list",
			Value: []any{[]any{float64(1.0), float64(2.0)}, []any{float64(3.0), float64(4.0)}},
		},
		{
			Field: "all_types.nested.id",
			Value: "nested_id",
		},
		{
			Field: "all_types.has_many",
			Value: map[string]any{
				"columns": []any{
					"id",
					"int",
					"float",
					"string",
					"bool",
					"timestamp",
					"int_list",
					"nested_int_pointer_list",
					"nested_int_list",
					"windowed_int__60__",
					"windowed_int__300__",
					"windowed_int__3600__",
					"dataclass",
					"dataclass_list",
				},
				"values": []any{
					// id column
					[]any{
						"id1",
						"id2",
					},
					// int column
					[]any{
						float64(123),
						float64(456),
					},
					// float column
					[]any{
						float64(123),
						float64(456),
					},
					// string column
					[]any{
						"abc",
						"def",
					},
					// bool column
					[]any{
						true,
						false,
					},
					// timestamp column
					[]any{
						"2024-05-09T22:29:00Z",
						"2024-05-09T22:30:00Z",
					},
					// int_list column
					[]any{
						[]any{float64(1), float64(2), float64(3)},
						[]any{float64(4), float64(5), float64(6)},
					},
					// nested_int_pointer_list column
					[]any{
						[]any{
							[]any{float64(1), float64(2)},
						},
						[]any{
							[]any{float64(3), float64(4)},
						},
					},
					// nested_int_list column
					[]any{
						[]any{
							[]any{float64(1), float64(2)},
						},
						[]any{
							[]any{float64(3), float64(4)},
						},
					},
					// windowed_int__60__ column
					[]any{
						float64(1),
						float64(2),
					},
					// windowed_int__300__ column
					[]any{
						float64(3),
						float64(4),
					},
					// windowed_int__3600__ column
					[]any{
						float64(5),
						float64(6),
					},
					// dataclass column
					[]any{
						[]any{float64(1.0), float64(2.0)},
						[]any{float64(3.0), float64(4.0)},
					},
					// dataclass_list column
					[]any{
						[]any{
							[]any{float64(1.0), float64(2.0)},
						},
						[]any{
							[]any{float64(3.0), float64(4.0)},
						},
					},
				},
			},
		},
	}
	result := OnlineQueryResult{
		Data:            data,
		Meta:            nil,
		features:        nil,
		expectedOutputs: nil,
	}
	features := allTypes{}
	unmarshalErr := result.UnmarshalInto(&features)
	if unmarshalErr != nil {
		t.Fatal(unmarshalErr)
	}
	assert.Nil(t, unmarshalErr)
	assert.Equal(t, int64(123), *features.Int)
	assert.Equal(t, float64(123), *features.Float)
	assert.Equal(t, "abc", *features.String)
	assert.Equal(t, true, *features.Bool)
	assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *features.Timestamp)
	assert.Equal(t, []int64{1, 2, 3}, *features.IntList)
	assert.Equal(t, 2, len(*features.NestedIntPointerList))
	assert.Equal(t, []int64{1, 2}, *(*features.NestedIntPointerList)[0])
	assert.Equal(t, []int64{3, 4}, *(*features.NestedIntPointerList)[1])
	assert.Equal(t, 2, len(*features.NestedIntList))
	assert.Equal(t, []int64{1, 2}, (*features.NestedIntList)[0])
	assert.Equal(t, []int64{3, 4}, (*features.NestedIntList)[1])
	assert.Equal(t, int64(1), *features.WindowedInt["1m"])
	assert.Equal(t, int64(2), *features.WindowedInt["5m"])
	assert.Equal(t, int64(3), *features.WindowedInt["1h"])
	assert.Equal(t, float64(1.0), *features.Dataclass.Lat)
	assert.Equal(t, float64(2.0), *features.Dataclass.Lng)
	assert.Equal(t, 2, len(*features.DataclassList))
	assert.Equal(t, float64(1.0), *(*features.DataclassList)[0].Lat)
	assert.Equal(t, float64(2.0), *(*features.DataclassList)[0].Lng)
	assert.Equal(t, float64(3.0), *(*features.DataclassList)[1].Lat)
	assert.Equal(t, float64(4.0), *(*features.DataclassList)[1].Lng)
	assert.Equal(t, "nested_id", *features.Nested.Id)
	assert.Nil(t, features.Nested.ShouldAlwaysBeNil)
	assert.Nil(t, features.Nested.Nested)

	// Has-manys
	assert.Equal(t, 2, len(*features.HasMany))
	assert.Equal(t, "id1", *(*features.HasMany)[0].Id)
	assert.Equal(t, "id2", *(*features.HasMany)[1].Id)
	assert.Equal(t, int64(123), *(*features.HasMany)[0].Int)
	assert.Equal(t, int64(456), *(*features.HasMany)[1].Int)
	assert.Equal(t, float64(123), *(*features.HasMany)[0].Float)
	assert.Equal(t, float64(456), *(*features.HasMany)[1].Float)
	assert.Equal(t, "abc", *(*features.HasMany)[0].String)
	assert.Equal(t, "def", *(*features.HasMany)[1].String)
	assert.Equal(t, true, *(*features.HasMany)[0].Bool)
	assert.Equal(t, false, *(*features.HasMany)[1].Bool)
	assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *(*features.HasMany)[0].Timestamp)
	assert.Equal(t, time.Date(2024, 5, 9, 22, 30, 0, 0, time.UTC), *(*features.HasMany)[1].Timestamp)
	assert.Equal(t, []int64{1, 2, 3}, *(*features.HasMany)[0].IntList)
	assert.Equal(t, []int64{4, 5, 6}, *(*features.HasMany)[1].IntList)
	assert.Equal(t, 1, len(*(*features.HasMany)[0].NestedIntPointerList))
	assert.Equal(t, 1, len(*(*features.HasMany)[1].NestedIntPointerList))
	assert.Equal(t, []int64{1, 2}, *(*(*features.HasMany)[0].NestedIntPointerList)[0])
	assert.Equal(t, []int64{3, 4}, *(*(*features.HasMany)[1].NestedIntPointerList)[0])
	assert.Equal(t, 1, len(*(*features.HasMany)[0].NestedIntList))
	assert.Equal(t, 1, len(*(*features.HasMany)[1].NestedIntList))
	assert.Equal(t, []int64{1, 2}, (*(*features.HasMany)[0].NestedIntList)[0])
	assert.Equal(t, []int64{3, 4}, (*(*features.HasMany)[1].NestedIntList)[0])
	assert.Equal(t, int64(1), *(*features.HasMany)[0].WindowedInt["1m"])
	assert.Equal(t, int64(2), *(*features.HasMany)[1].WindowedInt["1m"])
	assert.Equal(t, int64(3), *(*features.HasMany)[0].WindowedInt["5m"])
	assert.Equal(t, int64(4), *(*features.HasMany)[1].WindowedInt["5m"])
	assert.Equal(t, int64(5), *(*features.HasMany)[0].WindowedInt["1h"])
	assert.Equal(t, int64(6), *(*features.HasMany)[1].WindowedInt["1h"])
	assert.Equal(t, float64(1.0), *(*features.HasMany)[0].Dataclass.Lat)
	assert.Equal(t, float64(2.0), *(*features.HasMany)[0].Dataclass.Lng)
	assert.Equal(t, float64(3.0), *(*features.HasMany)[1].Dataclass.Lat)
	assert.Equal(t, float64(4.0), *(*features.HasMany)[1].Dataclass.Lng)
	assert.Equal(t, 1, len(*(*features.HasMany)[0].DataclassList))
	assert.Equal(t, float64(1.0), *(*(*features.HasMany)[0].DataclassList)[0].Lat)
	assert.Equal(t, float64(2.0), *(*(*features.HasMany)[0].DataclassList)[0].Lng)
	assert.Equal(t, 1, len(*(*features.HasMany)[1].DataclassList))
	assert.Equal(t, float64(3.0), *(*(*features.HasMany)[1].DataclassList)[0].Lat)
	assert.Equal(t, float64(4.0), *(*(*features.HasMany)[1].DataclassList)[0].Lng)
}

func TestUnmarshalVersionedFeatures(t *testing.T) {
	data := []FeatureResult{
		{
			Field:     "unmarshal_user.grade",
			Value:     1,
			Pkey:      "khjdsfjhdksjfh",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
		{
			Field:     "unmarshal_user.grade@2",
			Value:     2,
			Pkey:      "kjhsdfkjkdjfk",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
	}
	result := OnlineQueryResult{
		Data:            data,
		Meta:            nil,
		features:        nil,
		expectedOutputs: nil,
	}
	user := unmarshalUser{}
	unmarshalErr := result.UnmarshalInto(&user)
	if unmarshalErr != nil {
		t.Fatal(unmarshalErr)
	}
	assert.Nil(t, unmarshalErr)
	assert.Equal(t, 2, *user.Grade)
	assert.Equal(t, 1, *user.GradeV1)
	assert.Equal(t, 2, *user.GradeV2)
}

func TestUnmarshalWindowedFeatures(t *testing.T) {
	data := []FeatureResult{
		{
			Field:     "unmarshal_user.avg_spend__60__",
			Value:     60.0,
			Pkey:      "khjdsfjhdksjfh",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
		{
			Field:     "unmarshal_user.avg_spend__300__",
			Value:     300.0,
			Pkey:      "kjhsdfkjkdjfk",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
		{
			Field:     "unmarshal_user.avg_spend__3600__",
			Value:     3600.0,
			Pkey:      "kjhsdfkjkdjfk",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
	}
	result := OnlineQueryResult{
		Data:            data,
		Meta:            nil,
		features:        nil,
		expectedOutputs: nil,
	}
	user := unmarshalUser{}
	unmarshalErr := result.UnmarshalInto(&user)
	if unmarshalErr != nil {
		t.Fatal(unmarshalErr)
	}
	assert.Nil(t, unmarshalErr)
	assert.Equal(t, 60.0, *user.AvgSpend["1m"])
	assert.Equal(t, 300.0, *user.AvgSpend["5m"])
	assert.Equal(t, 3600.0, *user.AvgSpend["1h"])
}

func TestUnmarshalDataclassFeatures(t *testing.T) {
	data := []FeatureResult{
		{
			Field:     "unmarshal_user.lat_lng",
			Value:     []any{1.0, 2.0},
			Pkey:      "khjdsfjhdksjfh",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
	}
	result := OnlineQueryResult{
		Data:            data,
		Meta:            nil,
		features:        nil,
		expectedOutputs: nil,
	}
	user := unmarshalUser{}
	unmarshalErr := result.UnmarshalInto(&user)
	if unmarshalErr != nil {
		t.Fatal(unmarshalErr)
	}
	assert.Nil(t, unmarshalErr)
	assert.Equal(t, 1.0, *user.LatLng.Lat)
	assert.Equal(t, 2.0, *user.LatLng.Lng)
}

func TestUnmarshalWrongType(t *testing.T) {
	fqn := "unmarshal_user.int"
	data := []FeatureResult{
		{
			Field:     fqn,
			Value:     "1",
			Pkey:      "abc",
			Timestamp: time.Time{},
			Meta:      nil,
			Error:     nil,
		},
	}
	result := OnlineQueryResult{
		Data:            data,
		Meta:            nil,
		features:        nil,
		expectedOutputs: nil,
	}
	user := unmarshalUser{}
	unmarshalErr := result.UnmarshalInto(&user)
	if unmarshalErr == nil {
		fmt.Println("We successfully unmarshalled the wrong type into a struct field - the value is: ", *user.Int)
		t.Fatal("Expected an error when unmarshalling the wrong type into a struct field")
	} else {
		assert.Contains(t, unmarshalErr.Error(), fqn)
		assert.Contains(t, unmarshalErr.Error(), internal.KindMismatchError(reflect.Int64, reflect.String).Error())
		fmt.Println("We correctly surfaced an unmarshal type mismatch error - the error is: ", unmarshalErr)
	}
}

// Test primitives unmarshalling only.
func TestUnmarshalOnlineQueryBulkResultPrimitives(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.String: []string{"abc", "def"},
		testRootFeatures.AllTypes.Float:  []float64{1.0, 2.0},
		testRootFeatures.AllTypes.Bool:   []bool{true, false},
		testRootFeatures.AllTypes.Int:    []int{1, 2},
		testRootFeatures.AllTypes.Timestamp: []time.Time{
			time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC),
			time.Date(2024, 5, 9, 22, 30, 0, 0, time.UTC),
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(resultHolders))
	assert.Equal(t, "abc", *resultHolders[0].String)
	assert.Equal(t, "def", *resultHolders[1].String)
	assert.Equal(t, 1.0, *resultHolders[0].Float)
	assert.Equal(t, 2.0, *resultHolders[1].Float)
	assert.Equal(t, true, *resultHolders[0].Bool)
	assert.Equal(t, false, *resultHolders[1].Bool)
	assert.Equal(t, int64(1), *resultHolders[0].Int)
	assert.Equal(t, int64(2), *resultHolders[1].Int)
	assert.Equal(t, time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC), *resultHolders[0].Timestamp)
	assert.Equal(t, time.Date(2024, 5, 9, 22, 30, 0, 0, time.UTC), *resultHolders[1].Timestamp)
}

// TestUnmarshalOnlineQueryBulkResultDataclasses tests unmarshalling
// a bulk query result with dataclass features. This test first
// builds an Arrow table that contains a column of dataclass features,
// then unmarshals the table into appropriate structs.
func TestUnmarshalOnlineQueryBulkResultDataclasses(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	lat := 37.7749
	lng := 122.4194
	lat2 := 47.6062
	lng2 := 122.3321
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.Dataclass: []*testLatLng{
			{
				Lat: &lat,
				Lng: &lng,
			},
			nil,
			{
				Lat: &lat2,
				Lng: &lng2,
			},
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, len(resultHolders))
	assert.Equal(t, testLatLng{&lat, &lng}, *resultHolders[0].Dataclass)
	assert.Nil(t, resultHolders[1].Dataclass)
	assert.Equal(t, testLatLng{&lat2, &lng2}, *resultHolders[2].Dataclass)
}

// TestUnmarshalQueryBulkOptionalDataclassNested
func TestUnmarshalQueryBulkOptionalDataclassNested(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.DataclassWithDataclass: []*child{
			{
				Name: ptr.Ptr("Alice"),
				Mom: &parent{
					Name: ptr.Ptr("Alice's Mom"),
					Dad: &grandparent{
						Name: ptr.Ptr("Alice's Grandpa"),
					},
				},
			},
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(resultHolders))
	assert.Equal(t, "Alice", *resultHolders[0].DataclassWithDataclass.Name)
	assert.Equal(t, "Alice's Mom", *resultHolders[0].DataclassWithDataclass.Mom.Name)
	assert.Equal(t, "Alice's Grandpa", *resultHolders[0].DataclassWithDataclass.Mom.Dad.Name)
	assert.Nil(t, resultHolders[0].DataclassWithDataclass.Dad)
	assert.Nil(t, resultHolders[0].DataclassWithDataclass.Mom.Mom)
}

func TestUnmarshalBulkQueryDataclassWithOverrides(t *testing.T) {
	assert.Nil(t, InitFeatures(&testRootFeatures))
	name := "abc"
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.DataclassWithOverrides: []dclassWithOverrides{
			{
				CamelName: &name,
			},
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(resultHolders))
	assert.Equal(t, "abc", *resultHolders[0].DataclassWithOverrides.CamelName)
}

func TestUnmarshalBulkQueryDataclassList(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	lat1a := 37.7749
	lng1a := 122.4194
	lat1b := 47.6062
	lng1b := 122.3321
	lat2a := 43.6532
	lng2a := 79.3832
	lat2b := 40.7128
	lng2b := 74.0060

	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.DataclassList: [][]testLatLng{
			{
				{
					Lat: &lat1a,
					Lng: &lng1a,
				},
				{
					Lat: &lat1b,
					Lng: &lng1b,
				},
			},
			{
				{
					Lat: &lat2a,
					Lng: &lng2a,
				},
				{
					Lat: &lat2b,
					Lng: &lng2b,
				},
			},
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(resultHolders))
	assert.Equal(t, 2, len(*resultHolders[0].DataclassList))
	assert.Equal(t, testLatLng{&lat1a, &lng1a}, (*resultHolders[0].DataclassList)[0])
	assert.Equal(t, testLatLng{&lat1b, &lng1b}, (*resultHolders[0].DataclassList)[1])
	assert.Equal(t, 2, len(*resultHolders[1].DataclassList))
	assert.Equal(t, testLatLng{&lat2a, &lng2a}, (*resultHolders[1].DataclassList)[0])
	assert.Equal(t, testLatLng{&lat2b, &lng2b}, (*resultHolders[1].DataclassList)[1])
}

func TestUnmarshalBulkQueryNestedIntListWithInnerNilSlice(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.NestedIntList: [][][]int64{
			{
				{1, 2},
				nil,
				{3, 4},
			},
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(resultHolders))
	assert.Equal(t, 3, len(*resultHolders[0].NestedIntList))
	assert.Equal(t, []int64{1, 2}, (*resultHolders[0].NestedIntList)[0])
	assert.Equal(t, len((*resultHolders[0].NestedIntList)[1]), 0)
	assert.Equal(t, []int64{3, 4}, (*resultHolders[0].NestedIntList)[2])
}

func TestUnmarshalBulkQueryNestedIntPointerList(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.NestedIntPointerList: []*[]*[]int64{
			{
				{1, 2},
				{3, 4},
			},
			{
				{5, 6},
				{7, 8},
			},
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(resultHolders))
	assert.Equal(t, 2, len(*resultHolders[0].NestedIntPointerList))
	assert.Equal(t, []int64{1, 2}, *(*resultHolders[0].NestedIntPointerList)[0])
	assert.Equal(t, []int64{3, 4}, *(*resultHolders[0].NestedIntPointerList)[1])
	assert.Equal(t, 2, len(*resultHolders[1].NestedIntPointerList))
	assert.Equal(t, []int64{5, 6}, *(*resultHolders[1].NestedIntPointerList)[0])
	assert.Equal(t, []int64{7, 8}, *(*resultHolders[1].NestedIntPointerList)[1])
}

func TestUnmarshalBulkQueryNestedIntPointerListWithFirstLevelNil(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.NestedIntPointerList: []*[]*[]int64{
			{
				{1, 2},
				{3, 4},
			},
			nil,
			{
				{5, 6},
				{7, 8},
			},
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, len(resultHolders))
	assert.Equal(t, 2, len(*resultHolders[0].NestedIntPointerList))
	assert.Equal(t, []int64{1, 2}, *((*resultHolders[0].NestedIntPointerList)[0]))
	assert.Equal(t, []int64{3, 4}, *((*resultHolders[0].NestedIntPointerList)[1]))
	assert.Nil(t, resultHolders[1].NestedIntPointerList)
	assert.Equal(t, 2, len(*resultHolders[2].NestedIntPointerList))
	assert.Equal(t, []int64{5, 6}, *((*resultHolders[2].NestedIntPointerList)[0]))
	assert.Equal(t, []int64{7, 8}, *((*resultHolders[2].NestedIntPointerList)[1]))
}

func TestUnmarshalBulkQueryNestedPointerListWithInnerLevelNil(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.NestedIntPointerList: []*[]*[]int64{
			{
				{1, 2},
				nil,
				{3, 4},
			},
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(resultHolders))
	assert.Equal(t, 3, len(*resultHolders[0].NestedIntPointerList))
	assert.Equal(t, []int64{1, 2}, *(*resultHolders[0].NestedIntPointerList)[0])
	assert.Nil(t, (*resultHolders[0].NestedIntPointerList)[1])
	assert.Equal(t, []int64{3, 4}, *(*resultHolders[0].NestedIntPointerList)[2])
}

func TestUnmarshalBulkQueryDataclassWithList(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.DataclassWithList: []*favoriteThings{
			{
				Numbers: &[]int64{1, 2},
				Words:   &[]string{"abc", "def"},
			},
			{
				Numbers: &[]int64{3, 4},
				Words:   &[]string{"ghi", "jkl"},
			},
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(resultHolders))
	assert.Equal(t, 2, len(*resultHolders[0].DataclassWithList.Numbers))
	assert.Equal(t, int64(1), (*resultHolders[0].DataclassWithList.Numbers)[0])
	assert.Equal(t, int64(2), (*resultHolders[0].DataclassWithList.Numbers)[1])
	assert.Equal(t, 2, len(*resultHolders[0].DataclassWithList.Words))
	assert.Equal(t, "abc", (*resultHolders[0].DataclassWithList.Words)[0])
	assert.Equal(t, "def", (*resultHolders[0].DataclassWithList.Words)[1])

	assert.Equal(t, 2, len(*resultHolders[1].DataclassWithList.Numbers))
	assert.Equal(t, int64(3), (*resultHolders[1].DataclassWithList.Numbers)[0])
	assert.Equal(t, int64(4), (*resultHolders[1].DataclassWithList.Numbers)[1])
	assert.Equal(t, 2, len(*resultHolders[1].DataclassWithList.Words))
	assert.Equal(t, "ghi", (*resultHolders[1].DataclassWithList.Words)[0])
	assert.Equal(t, "jkl", (*resultHolders[1].DataclassWithList.Words)[1])
}

func TestUnmarshalBulkQueryDataclassWithNils(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.DataclassWithNils: []possessions{
			{
				Car:   ptr.Ptr("Toyota"),
				Yacht: nil,
				Plane: ptr.Ptr("Boeing"),
			},
			{
				Car:   ptr.Ptr("Honda"),
				Yacht: ptr.Ptr("Yamaha"),
				Plane: nil,
			},
		},
	}
	scalarsTable, scalarsErr := buildTableFromFeatureToValuesMap(scalarsMap)
	assert.Nil(t, scalarsErr)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: scalarsTable,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)

	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(resultHolders))
	assert.Equal(t, "Toyota", *resultHolders[0].DataclassWithNils.Car)
	assert.Nil(t, resultHolders[0].DataclassWithNils.Yacht)
	assert.Equal(t, "Boeing", *resultHolders[0].DataclassWithNils.Plane)

	assert.Equal(t, "Honda", *resultHolders[1].DataclassWithNils.Car)
	assert.Equal(t, "Yamaha", *resultHolders[1].DataclassWithNils.Yacht)
	assert.Nil(t, resultHolders[1].DataclassWithNils.Plane)
}

// TestUnmarshalBulkQueryOptionalValues tests that when a
// feature is optional, we can still unmarshal a bulk query
// result successfully.
func TestUnmarshalBulkQueryOptionalValues(t *testing.T) {
	schema := arrow.NewSchema([]arrow.Field{
		{Name: "all_types.string", Type: arrow.BinaryTypes.LargeString},
	}, nil)
	recordBuilder := array.NewRecordBuilder(
		memory.NewGoAllocator(),
		schema,
	)
	defer recordBuilder.Release()
	recordBuilder.Field(0).(*array.LargeStringBuilder).AppendValues(
		[]string{"abc", "def", "ghi"},
		[]bool{true, false, true},
	)
	table := array.NewTableFromRecords(schema, []arrow.Record{recordBuilder.NewRecord()})

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: table,
	}
	defer bulkRes.Release()

	resultHolders := make([]allTypes, 0)
	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, len(resultHolders))
	assert.Equal(t, "abc", *resultHolders[0].String)
	assert.Nil(t, resultHolders[1].String)
	assert.Equal(t, "ghi", *resultHolders[2].String)
}

// TestUnmarshalBulkQueryTimestampsWithUnitVariety tests that when features
// are timestamps, we correctly use the time unit to unmarshal the timestamps.
func TestUnmarshalBulkQueryTimestampsWithUnitVariety(t *testing.T) {
	for _, fixture := range []struct {
		unit           arrow.TimeUnit
		expectedTime   time.Time
		timestampValue int64
	}{
		{
			unit:           arrow.Microsecond,
			expectedTime:   time.Date(2024, 5, 9, 22, 29, 0, 111222000, time.UTC),
			timestampValue: time.Date(2024, 5, 9, 22, 29, 0, 111222333, time.UTC).UnixMicro(),
		},
		{
			unit:           arrow.Millisecond,
			expectedTime:   time.Date(2024, 5, 9, 22, 29, 0, 111000000, time.UTC),
			timestampValue: time.Date(2024, 5, 9, 22, 29, 0, 111222333, time.UTC).UnixMilli(),
		},
		{
			unit:           arrow.Second,
			expectedTime:   time.Date(2024, 5, 9, 22, 29, 0, 0, time.UTC),
			timestampValue: time.Date(2024, 5, 9, 22, 29, 0, 111222333, time.UTC).Unix(),
		},
		{
			unit:           arrow.Nanosecond,
			expectedTime:   time.Date(2024, 5, 9, 22, 29, 0, 111222333, time.UTC),
			timestampValue: time.Date(2024, 5, 9, 22, 29, 0, 111222333, time.UTC).UnixNano(),
		},
	} {
		t.Run(fmt.Sprintf("unit=%s", fixture.unit), func(t *testing.T) {
			schema := arrow.NewSchema([]arrow.Field{
				{Name: "all_types.timestamp", Type: &arrow.TimestampType{
					Unit:     fixture.unit,
					TimeZone: "UTC",
				}},
			}, nil)
			recordBuilder := array.NewRecordBuilder(
				memory.NewGoAllocator(),
				schema,
			)
			defer recordBuilder.Release()
			recordBuilder.Field(0).(*array.TimestampBuilder).AppendValues(
				[]arrow.Timestamp{
					arrow.Timestamp(fixture.timestampValue),
				}, nil,
			)
			table := array.NewTableFromRecords(schema, []arrow.Record{recordBuilder.NewRecord()})

			bulkRes := OnlineQueryBulkResult{
				ScalarsTable: table,
			}
			defer bulkRes.Release()

			resultHolders := make([]allTypes, 0)
			if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, 1, len(resultHolders))
			assert.Equal(t, fixture.expectedTime, *resultHolders[0].Timestamp)
		})
	}
}

// TestUnmarshalQueryBulkListOfPrimitives tests unmarshalling a column whose
// data type is a list of primitives.
func TestUnmarshalQueryBulkListOfPrimitives(t *testing.T) {
	// Not using `buildTableFromFeatureToValuesMap` like the scalar
	// primitives test above because `buildTableFromFeatureToValuesMap`
	// does not yet support converting a 2D list of primitives to an Arrow
	// array.
	// TODO: We can now use `buildTableFromFeatureToValuesMap` for this test.
	//       Migrate.
	encoded, readErr := os.ReadFile(filepath.Join(".", "internal", "sample_data", "list_of_primitives.txt"))
	if readErr != nil {
		log.Fatal(readErr)
	}

	bytesData, decodeErr := base64.StdEncoding.DecodeString(string(encoded))
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}

	bulkResponse := OnlineQueryBulkResponse{}
	assert.Nil(t, bulkResponse.Unmarshal(bytesData))
	singleResponse := bulkResponse.QueryResults["0"]
	assert.Equal(t, len(singleResponse.Errors), 0)

	bulkRes := OnlineQueryBulkResult{
		ScalarsTable: singleResponse.ScalarData,
		GroupsTables: singleResponse.GroupsData,
		Meta:         singleResponse.Meta,
	}
	defer bulkRes.Release()
	resultHolders := make([]user, 0)
	if err := bulkRes.UnmarshalInto(&resultHolders); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(resultHolders))

	assert.Equal(t, 3, len(*resultHolders[0].FavoriteNumbers))
	assert.Equal(t, int64(1), (*resultHolders[0].FavoriteNumbers)[0])
	assert.Equal(t, int64(2), (*resultHolders[0].FavoriteNumbers)[1])
	assert.Equal(t, int64(3), (*resultHolders[0].FavoriteNumbers)[2])

	assert.Equal(t, 3, len(*resultHolders[1].FavoriteColors))
	assert.Equal(t, "red", (*resultHolders[1].FavoriteColors)[0])
	assert.Equal(t, "green", (*resultHolders[1].FavoriteColors)[1])
	assert.Equal(t, "blue", (*resultHolders[1].FavoriteColors)[2])
}

func benchmarkUnmarshal(t *testing.T, data []FeatureResult, resultHolder any) time.Duration {
	t.Helper()
	result := OnlineQueryResult{
		Data:            data,
		Meta:            nil,
		features:        nil,
		expectedOutputs: nil,
	}
	var sum time.Duration
	var unmarshalErr error
	for i := 0; i < 100; i++ {
		start := time.Now()
		if unmarshalErr = result.UnmarshalInto(resultHolder); unmarshalErr != (*ClientError)(nil) {
			t.Fatal(unmarshalErr)
		}
		sum += time.Since(start)
	}
	return sum / 100
}

// TestEnsureTimelyUnmarshal tests that we maintain a non-quadratic unmarshal.
func TestEnsureTimelyUnmarshal(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)

	// Mimic JSON deser which returns all numbers as `float64`
	multiData := []FeatureResult{
		{
			Field: "all_types.int",
			Value: float64(123),
		},
		{
			Field: "all_types.float",
			Value: float64(123),
		},
		{
			Field: "all_types.string",
			Value: "abc",
		},
		{
			Field: "all_types.bool",
			Value: true,
		},
		{
			Field: "all_types.timestamp",
			Value: "2024-05-09T22:29:00Z",
		},
		{
			Field: "all_types.int_list",
			Value: []any{float64(1), float64(2), float64(3)},
		},
		{
			Field: "all_types.nested_int_pointer_list",
			Value: []any{[]any{float64(1), float64(2)}, []any{float64(3), float64(4)}},
		},
		{
			Field: "all_types.nested_int_list",
			Value: []any{[]any{float64(1), float64(2)}, []any{float64(3), float64(4)}},
		},
		{
			Field: "all_types.windowed_int__60__",
			Value: 1,
		},
		{
			Field: "all_types.windowed_int__300__",
			Value: 2,
		},
		{
			Field: "all_types.windowed_int__3600__",
			Value: 3,
		},
		{
			Field: "all_types.dataclass",
			Value: []any{float64(1.0), float64(2.0)},
		},
		{
			Field: "all_types.dataclass_list",
			Value: []any{[]any{float64(1.0), float64(2.0)}, []any{float64(3.0), float64(4.0)}},
		},
		{
			Field: "all_types.nested.id",
			Value: "nested_id",
		},
	}
	multiFeatures := &allTypes{}
	multiAvg := benchmarkUnmarshal(t, multiData, multiFeatures)

	singleData := []FeatureResult{
		{
			Field: "all_types.int",
			Value: float64(123),
		},
	}
	singleFeatures := &allTypes{}
	singleAvg := benchmarkUnmarshal(t, singleData, singleFeatures)
	assert.Equal(t, int64(123), *singleFeatures.Int)

	lenDelta := len(multiData) - len(singleData)
	delta := multiAvg - singleAvg
	deltaPerExtraItem := float64(delta) / float64(lenDelta)
	singleItemDuration := float64(singleAvg)
	multiplier := deltaPerExtraItem / singleItemDuration
	t.Logf(
		"multiplier (deltaPerExtraItem/singleItemDuration): %f, deltaPerExtraItem: %f, singleItemDuration: %f",
		multiplier, deltaPerExtraItem, singleItemDuration,
	)
	// Limit when run locally can be significantly less than 0.5,
	// but when run in CI it needs to be higher to not flake.
	limit := 0.5
	assert.True(t, multiplier < limit, "multiplier should be less than %v", limit)
}
