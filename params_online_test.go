package chalk

import (
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	assert "github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestOnlineQueryParamsAllTypes(t *testing.T) {
	// Tests that all types of input, output, and staleness parameters can be passed
	// without error.
	initErr := InitFeatures(&testRootFeatures)
	assert.Nil(t, initErr)
	params := OnlineQueryParams{}.
		WithInput(testRootFeatures.AllTypes.String, 1).
		WithOutputs(
			"all_types.string",
			testRootFeatures.AllTypes.Float,
			testRootFeatures.AllTypes.Int,
			testRootFeatures.AllTypes.Timestamp,
			testRootFeatures.AllTypes.IntList,
			testRootFeatures.AllTypes.WindowedInt,
			testRootFeatures.AllTypes.WindowedInt["1m"],
			testRootFeatures.AllTypes.WindowedInt["5m"],
			testRootFeatures.AllTypes.WindowedInt["1h"],
			testRootFeatures.AllTypes.WindowedList,
			testRootFeatures.AllTypes.WindowedList["1m"],
			testRootFeatures.AllTypes.Nested,
			testRootFeatures.AllTypes.Nested.Id,
			testRootFeatures.AllTypes.Dataclass,
			testRootFeatures.AllTypes.Dataclass.Lat,
			testRootFeatures.AllTypes.Dataclass.Lng,
		).
		WithStaleness(
			testRootFeatures.AllTypes.Bool, time.Second*5,
		)
	assert.Empty(t, params.underlying.builderErrors)
}

func TestOnlineQueryInputParamInteger(t *testing.T) {
	// Tests passing an integer as input feature reference. Should fail.
	var invalidFeatureReference int
	params := OnlineQueryParams{}.WithInput(invalidFeatureReference, "1")
	assert.NotEmpty(t, params.underlying.builderErrors)
}

func TestOnlineQueryOutputParamInteger(t *testing.T) {
	// Tests passing an integer as output feature reference. Should fail.
	var invalidFeatureReference int
	params := OnlineQueryParams{}.WithOutputs(invalidFeatureReference)
	assert.NotEmpty(t, params.underlying.builderErrors)
}

func TestOnlineQueryStalenessParamInteger(t *testing.T) {
	// Tests passing an integer as staleness feature reference. Should fail.
	var invalidFeatureReference int
	underlying := OnlineQueryParams{}.WithStaleness(invalidFeatureReference, time.Second*5)
	assert.NotEmpty(t, underlying.builderErrors)
}

// Tests that OnlineQuery successfully serializes all types of input feature values.
func TestOnlineQueryInputsAllTypes(t *testing.T) {
	t.Parallel()
	assert.Nil(t, InitFeatures(&testRootFeatures))
	timestamp := time.Date(2021, 1, 2, 3, 4, 45, 123, time.UTC)
	params := OnlineQueryParams{}.
		WithInput(testRootFeatures.AllTypes.HasMany, []hasMany{
			{
				Id:        lo.ToPtr("1"),
				Int:       lo.ToPtr(int64(1)),
				Float:     lo.ToPtr(float64(1.1)),
				String:    lo.ToPtr("abc"),
				Bool:      lo.ToPtr(true),
				Timestamp: &timestamp,
				IntList:   &[]int64{1, 2, 3},
				NestedIntPointerList: &[]*[]int64{
					{int64(1), int64(2)},
				},
				NestedIntList: &[][]int64{
					{int64(1), int64(2)},
				},
				WindowedInt: map[string]*int64{
					"1m": lo.ToPtr(int64(1)),
					"5m": lo.ToPtr(int64(2)),
					"1h": lo.ToPtr(int64(3)),
				},
				WindowedList: map[string]*[]int64{
					"1m": {1, 2, 3},
					"5m": {4, 5, 6},
					"1h": {7, 8, 9},
				},
				Dataclass: &testLatLng{
					Lat: lo.ToPtr(1.1),
					Lng: lo.ToPtr(2.2),
				},
				DataclassList: &[]testLatLng{
					{
						Lat: lo.ToPtr(3.3),
						Lng: lo.ToPtr(4.4),
					},
				},
				DataclassWithList: &favoriteThings{
					Numbers: &[]int64{1, 2, 3},
					Words:   &[]string{"a", "b", "c"},
				},
				DataclassWithNils: &possessions{
					Car:   lo.ToPtr("car"),
					Yacht: lo.ToPtr("yacht"),
					Plane: lo.ToPtr("plane"),
				},
				DataclassWithDataclass: &child{
					Name: lo.ToPtr("child"),
					Mom: &parent{
						Name: lo.ToPtr("mom"),
						Mom: &grandparent{
							Name: lo.ToPtr("grandma"),
						},
					},
				},
				DataclassWithOverrides: &dclassWithOverrides{
					CamelName: lo.ToPtr("camel"),
				},
			},
			{
				Id:        lo.ToPtr("2"),
				Int:       lo.ToPtr(int64(2)),
				Float:     lo.ToPtr(float64(2.2)),
				String:    lo.ToPtr("def"),
				Bool:      lo.ToPtr(false),
				Timestamp: &timestamp,
				IntList:   &[]int64{4, 5, 6},
				NestedIntPointerList: &[]*[]int64{
					{int64(3), int64(4)},
				},
				NestedIntList: &[][]int64{
					{int64(3), int64(4)},
				},
				WindowedInt: map[string]*int64{
					"1m": lo.ToPtr(int64(4)),
					"5m": lo.ToPtr(int64(5)),
					"1h": lo.ToPtr(int64(6)),
				},
				WindowedList: map[string]*[]int64{
					"1m": {4, 5, 6},
					"5m": {7, 8, 9},
					"1h": {10, 11, 12},
				},
				Dataclass: &testLatLng{
					Lat: lo.ToPtr(5.5),
					Lng: lo.ToPtr(6.6),
				},
				DataclassList: &[]testLatLng{
					{
						Lat: lo.ToPtr(7.7),
						Lng: lo.ToPtr(8.8),
					},
				},
				DataclassWithList: &favoriteThings{
					Numbers: &[]int64{4, 5, 6},
					Words:   &[]string{"d", "e", "f"},
				},
				DataclassWithNils: &possessions{
					Car:   lo.ToPtr("car2"),
					Yacht: lo.ToPtr("yacht2"),
					Plane: lo.ToPtr("plane2"),
				},
				DataclassWithDataclass: &child{
					Name: lo.ToPtr("child2"),
					Mom: &parent{
						Name: lo.ToPtr("mom2"),
					},
				},
				DataclassWithOverrides: &dclassWithOverrides{
					CamelName: lo.ToPtr("camel2"),
				},
			},
		})
	request, err := params.underlying.serialize()
	assert.NoError(t, err)
	assert.NotNil(t, request)

	path := filepath.Join("internal", "fixtures", "online_query_inputs_all_types.json")

	// Read the entire file content into a byte slice
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	inputBytes, err := json.MarshalIndent(request.Inputs, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, string(fileContent), string(inputBytes))

	err = os.WriteFile(path, inputBytes, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
