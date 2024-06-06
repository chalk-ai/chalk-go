package chalk

import (
	"encoding/base64"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	assert "github.com/stretchr/testify/require"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type unmarshalLatLng struct {
	Lat *float64 `dataclass_field:"true"`
	Lng *float64 `dataclass_field:"true"`
}

type unmarshalUser struct {
	Id *string

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
	data := []FeatureResult{
		{
			Field:     "unmarshal_user.id",
			Value:     1,
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
		fmt.Println("We successfully unmarshalled the wrong type into a struct field - the value is: ", *user.Id)
		t.Fatal("Expected an error when unmarshalling the wrong type into a struct field")
	} else {
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

// TestUnmarshalOnlineQueryBulkResultDataclasses does what it says it does.
func TestUnmarshalOnlineQueryBulkResultDataclasses(t *testing.T) {
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	coord := 1.09
	scalarsMap := map[any]any{
		testRootFeatures.AllTypes.Dataclass: []testLatLng{
			{
				Lat: &coord,
				Lng: &coord,
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
	assert.Equal(t, testLatLng{&coord, &coord}, *resultHolders[0].Dataclass)
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
