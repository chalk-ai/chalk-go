package chalk

import (
	"encoding/json"
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	assert "github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var initErr error

func init() {
	initErr = InitFeatures(&testRootFeatures)
}

func TestOnlineQueryParamsAllTypes(t *testing.T) {
	// Tests that all types of input, output, and staleness parameters can be passed
	// without error.
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

// Tests that Feature structs have their nil fields omitted by default,
// and not omitted when `chalk:"dontomit"` flag is set.
func TestOnlineQueryParamsOmitNilFields(t *testing.T) {
	t.Parallel()
	assert.Nil(t, initErr)
	params := OnlineQueryParams{}.
		WithInput(testRootFeatures.AllTypes.Nested, levelOneNest{
			Id: ptr.Ptr("1"),
		}).
		WithInput(testRootFeatures.AllTypes.Dataclass, testLatLng{
			Lat: ptr.Ptr(1.1),
		})
	request, err := params.underlying.serialize()
	assert.NoError(t, err)
	assert.NotNil(t, request)

	path := filepath.Join("internal", "fixtures", "online_query_params_omit_nil_fields.json")
	fileContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	inputJsonBytes, err := json.MarshalIndent(request.Inputs, "", "  ")
	assert.NoError(t, err)
	err = os.WriteFile(path, inputJsonBytes, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	assert.Equal(t, string(fileContent), string(inputJsonBytes))

	bulkInputs, err := internal.SingleInputsToBulkInputs(params.underlying.inputs)
	assert.NoError(t, err)
	arrowBytes, err := internal.InputsToArrowBytes(bulkInputs)
	assert.NoError(t, err)
	assert.NotNil(t, arrowBytes)

	table, err := internal.ConvertBytesToTable(arrowBytes)
	assert.NoError(t, err)
	assert.NotNil(t, table)

	rows, err := internal.ExtractFeaturesFromTable(table, false)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(rows))

	featherInputJsonBytes, err := json.MarshalIndent(rows[0], "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, string(fileContent), string(featherInputJsonBytes))
}

// Tests that OnlineQuery successfully serializes all types of input feature values.
func TestOnlineQueryInputsAllTypes(t *testing.T) {
	t.Parallel()
	assert.Nil(t, initErr)
	timestamp := time.Date(2021, 1, 2, 3, 4, 45, 123, time.UTC)
	params := OnlineQueryParams{}.
		WithInput(testRootFeatures.AllTypes.Nested, levelOneNest{
			Id: ptr.Ptr("1"),
			Nested: &levelTwoNest{
				Id: ptr.Ptr("2"),
			},
		}).
		WithInput(testRootFeatures.AllTypes.HasMany, []hasMany{
			{
				Id:        ptr.Ptr("1"),
				Int:       ptr.Ptr(int64(1)),
				Float:     ptr.Ptr(float64(1.1)),
				String:    ptr.Ptr("abc"),
				Bool:      ptr.Ptr(true),
				Timestamp: &timestamp,
				IntList:   &[]int64{1, 2, 3},
				NestedIntPointerList: &[]*[]int64{
					{int64(1), int64(2)},
				},
				NestedIntList: &[][]int64{
					{int64(1), int64(2)},
				},
				WindowedInt: map[string]*int64{
					"1m": ptr.Ptr(int64(1)),
					"5m": ptr.Ptr(int64(2)),
					"1h": ptr.Ptr(int64(3)),
				},
				WindowedList: map[string]*[]int64{
					"1m": {1, 2, 3},
					"5m": {4, 5, 6},
					"1h": {7, 8, 9},
				},
				Dataclass: &testLatLng{
					Lat: ptr.Ptr(1.1),
					Lng: ptr.Ptr(2.2),
				},
				DataclassList: &[]testLatLng{
					{
						Lat: ptr.Ptr(3.3),
						Lng: ptr.Ptr(4.4),
					},
				},
				DataclassWithList: &favoriteThings{
					Numbers: &[]int64{1, 2, 3},
					Words:   &[]string{"a", "b", "c"},
				},
				DataclassWithNils: &possessions{
					Car:   ptr.Ptr("car"),
					Yacht: ptr.Ptr("yacht"),
					Plane: ptr.Ptr("plane"),
				},
				DataclassWithDataclass: &child{
					Name: ptr.Ptr("child"),
					Mom: &parent{
						Name: ptr.Ptr("mom"),
						Mom: &grandparent{
							Name: ptr.Ptr("grandma"),
						},
					},
				},
				DataclassWithOverrides: &dclassWithOverrides{
					CamelName: ptr.Ptr("camel"),
				},
			},
			{
				Id:        ptr.Ptr("2"),
				Int:       ptr.Ptr(int64(2)),
				Float:     ptr.Ptr(float64(2.2)),
				String:    ptr.Ptr("def"),
				Bool:      ptr.Ptr(false),
				Timestamp: &timestamp,
				IntList:   &[]int64{4, 5, 6},
				NestedIntPointerList: &[]*[]int64{
					{int64(3), int64(4)},
				},
				NestedIntList: &[][]int64{
					{int64(3), int64(4)},
				},
				WindowedInt: map[string]*int64{
					"1m": ptr.Ptr(int64(4)),
					"5m": ptr.Ptr(int64(5)),
					"1h": ptr.Ptr(int64(6)),
				},
				WindowedList: map[string]*[]int64{
					"1m": {4, 5, 6},
					"5m": {7, 8, 9},
					"1h": {10, 11, 12},
				},
				Dataclass: &testLatLng{
					Lat: ptr.Ptr(5.5),
					Lng: ptr.Ptr(6.6),
				},
				DataclassList: &[]testLatLng{
					{
						Lat: ptr.Ptr(7.7),
						Lng: ptr.Ptr(8.8),
					},
				},
				DataclassWithList: &favoriteThings{
					Numbers: &[]int64{4, 5, 6},
					Words:   &[]string{"d", "e", "f"},
				},
				DataclassWithNils: &possessions{
					Car:   ptr.Ptr("car2"),
					Yacht: ptr.Ptr("yacht2"),
					Plane: ptr.Ptr("plane2"),
				},
				DataclassWithDataclass: &child{
					Name: ptr.Ptr("child2"),
					Mom: &parent{
						Name: ptr.Ptr("mom2"),
					},
				},
				DataclassWithOverrides: &dclassWithOverrides{
					CamelName: ptr.Ptr("camel2"),
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
	err = os.WriteFile(path, inputBytes, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	assert.Equal(t, string(fileContent), string(inputBytes))
}

func TestWithInputsMapFromOnlineQueryParams(t *testing.T) {
	t.Parallel()
	assert.Nil(t, initErr)

	inputs := map[any]any{
		testRootFeatures.AllTypes.String: 1,
		testRootFeatures.AllTypes.Float:  1.1,
	}
	params := OnlineQueryParams{}.WithInputs(inputs)

	feature1, err := UnwrapFeature(testRootFeatures.AllTypes.String)
	assert.Nil(t, err)
	feature2, err := UnwrapFeature(testRootFeatures.AllTypes.Float)
	assert.Nil(t, err)

	_, ok := params.underlying.inputs[feature1.Fqn]
	assert.True(t, ok)
	_, ok = params.underlying.inputs[feature2.Fqn]
	assert.True(t, ok)
}

func TestWithInputsMapFromOnlineQueryParamsWithInputs(t *testing.T) {
	t.Parallel()
	assert.Nil(t, initErr)

	inputs := map[any]any{
		testRootFeatures.AllTypes.String: 1,
		testRootFeatures.AllTypes.Float:  1.1,
	}
	params := OnlineQueryParams{}.WithInput(testRootFeatures.AllTypes.Bool, true).WithInputs(inputs)

	feature1, err := UnwrapFeature(testRootFeatures.AllTypes.String)
	assert.Nil(t, err)
	feature2, err := UnwrapFeature(testRootFeatures.AllTypes.Float)
	assert.Nil(t, err)

	_, ok := params.underlying.inputs[feature1.Fqn]
	assert.True(t, ok)
	_, ok = params.underlying.inputs[feature2.Fqn]
	assert.True(t, ok)
}

func TestWithInputsMapFromOnlineQueryParamsWithOutputs(t *testing.T) {
	t.Parallel()
	assert.Nil(t, initErr)

	inputs := map[any]any{
		testRootFeatures.AllTypes.String: 1,
		testRootFeatures.AllTypes.Float:  1.1,
	}
	params := OnlineQueryParams{}.WithOutputs(testRootFeatures.AllTypes.Bool).WithInputs(inputs)

	feature1, err := UnwrapFeature(testRootFeatures.AllTypes.String)
	assert.Nil(t, err)
	feature2, err := UnwrapFeature(testRootFeatures.AllTypes.Float)
	assert.Nil(t, err)

	_, ok := params.underlying.inputs[feature1.Fqn]
	assert.True(t, ok)
	_, ok = params.underlying.inputs[feature2.Fqn]
	assert.True(t, ok)
}

func TestWithInputsMapFromOnlineQueryParamsComplete(t *testing.T) {
	t.Parallel()
	assert.Nil(t, initErr)

	inputs := map[any]any{
		testRootFeatures.AllTypes.String: 1,
		testRootFeatures.AllTypes.Float:  1.1,
	}
	params := OnlineQueryParamsComplete{}.WithInputs(inputs)

	feature1, err := UnwrapFeature(testRootFeatures.AllTypes.String)
	assert.Nil(t, err)
	feature2, err := UnwrapFeature(testRootFeatures.AllTypes.Float)
	assert.Nil(t, err)

	_, ok := params.underlying.inputs[feature1.Fqn]
	assert.True(t, ok)
	_, ok = params.underlying.inputs[feature2.Fqn]
	assert.True(t, ok)
}

func TestNamedQuery(t *testing.T) {
	t.Parallel()
	queryName := "test_query_name"
	var allParams []OnlineQueryParamsComplete
	allParams = append(allParams, OnlineQueryParams{}.WithInput("user.id", 1).WithQueryName(queryName))
	allParams = append(allParams, OnlineQueryParamsComplete{}.WithQueryName(queryName).WithInput("user.id", 1))
	allParams = append(allParams, OnlineQueryParamsComplete{}.WithBranchId("branch-1").WithInput("user.id", 1).WithQueryName(queryName))

	for _, params := range allParams {
		assert.Equal(t, queryName, params.underlying.QueryName)
	}
}
