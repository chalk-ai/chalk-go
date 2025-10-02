package chalk

import (
	"encoding/json"
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	assert "github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestOnlineQueryParamsAllTypes(t *testing.T) {
	t.Parallel()
	// Tests that all types of input, output, and staleness parameters can be passed
	// without error.
	fixtureRoot, initErr := GetRootFeatures()
	assert.NoError(t, initErr)
	params := OnlineQueryParams{}.
		WithInput(fixtureRoot.AllTypes.String, 1).
		WithOutputs(
			"all_types.string",
			fixtureRoot.AllTypes.Float,
			fixtureRoot.AllTypes.Int,
			fixtureRoot.AllTypes.Timestamp,
			fixtureRoot.AllTypes.IntList,
			fixtureRoot.AllTypes.WindowedInt,
			fixtureRoot.AllTypes.WindowedInt["1m"],
			fixtureRoot.AllTypes.WindowedInt["5m"],
			fixtureRoot.AllTypes.WindowedInt["1h"],
			fixtureRoot.AllTypes.WindowedList,
			fixtureRoot.AllTypes.WindowedList["1m"],
			fixtureRoot.AllTypes.Nested,
			fixtureRoot.AllTypes.Nested.Id,
			fixtureRoot.AllTypes.Dataclass,
			fixtureRoot.AllTypes.Dataclass.Lat,
			fixtureRoot.AllTypes.Dataclass.Lng,
		).
		WithStaleness(
			fixtureRoot.AllTypes.Bool, time.Second*5,
		)
	_, err := params.underlying.resolveSingle()
	assert.NoError(t, err)
}

func TestOnlineQueryInputParamInteger(t *testing.T) {
	t.Parallel()
	// Tests passing an integer as input feature reference. Should fail.
	var invalidFeatureReference int
	params := OnlineQueryParams{}.WithInput(invalidFeatureReference, "1")
	_, err := params.underlying.resolveSingle()
	assert.Error(t, err)
}

func TestOnlineQueryOutputParamInteger(t *testing.T) {
	t.Parallel()
	// Tests passing an integer as output feature reference. Should fail.
	var invalidFeatureReference int
	params := OnlineQueryParams{}.WithOutputs(invalidFeatureReference)
	_, err := params.underlying.resolveSingle()
	assert.Error(t, err)
}

func TestOnlineQueryStalenessParamInteger(t *testing.T) {
	t.Parallel()
	// Tests passing an integer as staleness feature reference. Should fail.
	var invalidFeatureReference int
	underlying := OnlineQueryParams{}.WithStaleness(invalidFeatureReference, time.Second*5)
	_, err := underlying.resolveSingle()
	assert.Error(t, err)
}

// Tests that Feature structs have their nil fields omitted by default,
// and not omitted when `chalk:"dontomit"` flag is set.
func TestOnlineQueryParamsOmitNilFields(t *testing.T) {
	t.Parallel()
	fixtureRoot, initErr := GetRootFeatures()
	assert.NoError(t, initErr)
	params := OnlineQueryParams{}.
		WithInput(fixtureRoot.AllTypes.Nested, fixtures.LevelOneNest{
			Id: ptr.New("1"),
		}).
		WithInput(fixtureRoot.AllTypes.Dataclass, fixtures.LatLng{
			Lat: ptr.New(1.1),
		})

	resolved, err := params.underlying.resolveSingle()
	assert.NoError(t, err)

	request, err := serializeOnlineQueryParams(&params.underlying, resolved)
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

	bulkInputs, err := internal.SingleInputsToBulkInputs(resolved.inputs)
	assert.NoError(t, err)
	arrowBytes, err := internal.InputsToArrowBytes(bulkInputs, fixtures.TestAllocator)
	assert.NoError(t, err)
	assert.NotNil(t, arrowBytes)

	table, err := internal.ConvertBytesToTable(arrowBytes, fixtures.TestAllocator)
	assert.NoError(t, err)
	assert.NotNil(t, table)

	rows, _, err := internal.ExtractFeaturesFromTable(table, false)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(rows))

	featherInputJsonBytes, err := json.MarshalIndent(rows[0], "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, string(fileContent), string(featherInputJsonBytes))
}

// TestBulkInputsOmitNilFields tests that for bulk inputs whose
// feature is nil for every row in the input, the feature is
// omitted from the serialized input (Arrow table) unless the
// feature has the `chalk:"dontomit"` tag. This also tests that
// dataclass fields are never omitted.
func TestBulkInputsOmitNilFields(t *testing.T) {
	t.Parallel()

	type omitTxn struct {
		Id       *string
		Amount   *int
		Cashback *int
	}

	type omitHasManyRoot struct {
		Id   *string
		Txns []*omitTxn
	}

	type omitHasOneRoot struct {
		Id  *string
		Txn *omitTxn
	}

	type omitDont struct {
		Id       *string
		Amount   *int
		Cashback *int `chalk:"dontomit"`
	}

	type omitDataclass struct {
		Id       *string `dataclass_field:"true"`
		Amount   *int
		Cashback *int
	}

	root := filepath.Join("internal", "fixtures", "field_omission")
	for _, fixture := range []struct {
		name     string
		input    map[string]any
		filename string
	}{
		{
			name:     "has-many inter-row",
			filename: "has_many_inter_row.json",
			input: map[string]any{
				"user.id":   []string{"user_1", "user_2", "user_3"},
				"user.name": []string{"Alice", "Bob", "Chinedum"},
				"user.txns": [][]omitTxn{
					{{Id: ptr.New("txn_1"), Amount: ptr.New(100)}},
					{},
					{{Id: ptr.New("txn_3")}},
				},
			},
		},
		{
			name:     "has-many intra-row",
			filename: "has_many_intra_row.json",
			input: map[string]any{
				"user.id": []string{"user_1"},
				"user.txns": [][]omitTxn{
					{
						{Id: ptr.New("txn_1"), Amount: ptr.New(100)},
						{Id: ptr.New("txn_2")},
					},
				},
			},
		},
		{
			name:     "has-one",
			filename: "has_one.json",
			input: map[string]any{
				"user.id": []string{"user_1", "user_2", "user_3"},
				"user.txn": []omitTxn{
					{Id: ptr.New("txn_1"), Amount: ptr.New(100)},
					{},
					{Id: ptr.New("txn_3")},
				},
			},
		},
		{
			name:     "has-one with optional",
			filename: "has_one_with_optional.json",
			input: map[string]any{
				"user.id": []string{"user_1", "user_2", "user_3"},
				"user.txn": []*omitTxn{
					{Id: ptr.New("txn_1"), Amount: ptr.New(100)},
					nil,
					{Id: ptr.New("txn_3")},
				},
			},
		},
		{
			name:     "has-one -> has-one",
			filename: "has_one_has_one.json",
			input: map[string]any{
				"user.id": []string{"user_1", "user_2", "user_3"},
				"user.has_one": []omitHasOneRoot{
					{
						Id:  ptr.New("root_1"),
						Txn: &omitTxn{Id: ptr.New("txn_1"), Amount: ptr.New(100)},
					},
					{},
					{
						Id:  ptr.New("root_3"),
						Txn: &omitTxn{Id: ptr.New("txn_3")},
					},
				},
			},
		},
		{
			name:     "has-many -> has-one",
			filename: "has_many_has_one.json",
			input: map[string]any{
				"user.id": []string{"user_1", "user_2", "user_3"},
				"user.has_many": [][]omitHasOneRoot{
					{
						{
							Id:  ptr.New("root_1"),
							Txn: &omitTxn{Id: ptr.New("txn_1"), Amount: ptr.New(100)},
						},
					},
					{},
					{
						{
							Id:  ptr.New("root_3"),
							Txn: &omitTxn{Id: ptr.New("txn_3")},
						},
					},
				},
			},
		},
		{
			name:     "has-one -> has-many",
			filename: "has_one_has_many.json",
			input: map[string]any{
				"user.id": []string{"user_1", "user_2", "user_3"},
				"user.has_one": []omitHasManyRoot{
					{
						Id: ptr.New("root_1"),
						Txns: []*omitTxn{
							{Id: ptr.New("txn_1"), Amount: ptr.New(100)},
						},
					},
					{},
					{
						Id: ptr.New("root_3"),
						Txns: []*omitTxn{
							{Id: ptr.New("txn_3")},
						},
					},
				},
			},
		},
		{
			name:     "has-many -> has-many",
			filename: "has_many_has_many.json",
			input: map[string]any{
				"user.id": []string{"user_1", "user_2", "user_3"},
				"user.has_many": [][]omitHasManyRoot{
					{
						{
							Id: ptr.New("root_1"),
							Txns: []*omitTxn{
								{Id: ptr.New("txn_1"), Amount: ptr.New(100)},
							},
						},
					},
					{},
					{
						{
							Id: ptr.New("root_3"),
							Txns: []*omitTxn{
								{Id: ptr.New("txn_3")},
							},
						},
					},
				},
			},
		},
		{
			name:     "dataclass",
			filename: "dataclass.json",
			input: map[string]any{
				"user.id": []string{"user_1", "user_2", "user_3"},
				"user.dataclass": []omitDataclass{
					{Id: ptr.New("txn_1"), Amount: ptr.New(100)},
					{},
					{Id: ptr.New("txn_3")},
				},
			},
		},
		{
			name:     "dataclass with nil",
			filename: "dataclass_with_nil.json",
			input: map[string]any{
				"user.id": []string{"user_1", "user_2", "user_3"},
				"user.dataclass": []*omitDataclass{
					{Id: ptr.New("txn_1"), Amount: ptr.New(100)},
					nil,
					{Id: ptr.New("txn_3")},
				},
			},
		},
		{
			name:     "list of dataclass",
			filename: "list_of_dataclass.json",
			input: map[string]any{
				"user.id": []string{"user_1", "user_2", "user_3"},
				"user.dataclasses": [][]omitDataclass{
					{{Id: ptr.New("txn_1"), Amount: ptr.New(100)}},
					{},
					{{Id: ptr.New("txn_3")}},
				},
			},
		},
		{
			name:     "dontomit tag",
			filename: "dont_omit.json",
			input: map[string]any{
				"user.id": []string{"user_1", "user_2", "user_3"},
				"user.dont": []omitDont{
					{Id: ptr.New("txn_1"), Amount: ptr.New(100)},
					{},
					{Id: ptr.New("txn_3")},
				},
			},
		},
	} {
		t.Run(fixture.name, func(t *testing.T) {
			t.Parallel()
			table, err := tableFromFqnToValues(fixture.input)
			assert.NoError(t, err)

			rows, _, err := internal.ExtractFeaturesFromTable(table, false)
			assert.NoError(t, err)

			featherInputJsonBytes, err := json.MarshalIndent(rows, "", "  ")
			assert.NoError(t, err)

			fileContent, err := os.ReadFile(filepath.Join(root, fixture.filename))
			if err != nil {
				fileContent = []byte("")
			}

			//assert.NoError(t, os.WriteFile(filepath.Join(root, fixture.filename), featherInputJsonBytes, 0644))
			assert.Equal(t, string(fileContent), string(featherInputJsonBytes))
		})
	}

}

// Tests that OnlineQuery successfully serializes all types of input feature values.
func TestOnlineQueryInputsAllTypes(t *testing.T) {
	t.Parallel()
	fixtureRoot, initErr := GetRootFeatures()
	assert.NoError(t, initErr)
	timestamp := time.Date(2021, 1, 2, 3, 4, 45, 123, time.UTC)
	params := OnlineQueryParams{}.
		WithInput(fixtureRoot.AllTypes.Nested, fixtures.LevelOneNest{
			Id: ptr.New("1"),
			Nested: &fixtures.LevelTwoNest{
				Id: ptr.New("2"),
			},
		}).
		WithInput(fixtureRoot.AllTypes.HasMany, []fixtures.HasMany{
			{
				Id:        ptr.New("1"),
				Int:       ptr.New(int64(1)),
				Float:     ptr.New(float64(1.1)),
				String:    ptr.New("abc"),
				Bool:      ptr.New(true),
				Timestamp: &timestamp,
				IntList:   &[]int64{1, 2, 3},
				NestedIntPointerList: &[]*[]int64{
					{int64(1), int64(2)},
				},
				NestedIntList: &[][]int64{
					{int64(1), int64(2)},
				},
				WindowedInt: map[string]*int64{
					"1m": ptr.New(int64(1)),
					"5m": ptr.New(int64(2)),
					"1h": ptr.New(int64(3)),
				},
				WindowedList: map[string]*[]int64{
					"1m": {1, 2, 3},
					"5m": {4, 5, 6},
					"1h": {7, 8, 9},
				},
				Dataclass: &fixtures.LatLng{
					Lat: ptr.New(1.1),
					Lng: ptr.New(2.2),
				},
				DataclassList: &[]fixtures.LatLng{
					{
						Lat: ptr.New(3.3),
						Lng: ptr.New(4.4),
					},
				},
				DataclassWithList: &fixtures.FavoriteThings{
					Numbers: &[]int64{1, 2, 3},
					Words:   &[]string{"a", "b", "c"},
				},
				DataclassWithNils: &fixtures.Possessions{
					Car:   ptr.New("car"),
					Yacht: ptr.New("yacht"),
					Plane: ptr.New("plane"),
				},
				// CHA-5430
				//DataclassWithDataclass: &fixtures.Child{
				//	Name: ptr.New("fixtures.Child"),
				//	Mom: &fixtures.Parent{
				//		Name: ptr.New("mom"),
				//		Mom: &fixtures.Grandparent{
				//			Name: ptr.New("grandma"),
				//		},
				//	},
				//},
				DataclassWithOverrides: &fixtures.DclassWithOverrides{
					CamelName: ptr.New("camel"),
				},
			},
			{
				Id:        ptr.New("2"),
				Int:       ptr.New(int64(2)),
				Float:     ptr.New(float64(2.2)),
				String:    ptr.New("def"),
				Bool:      ptr.New(false),
				Timestamp: &timestamp,
				IntList:   &[]int64{4, 5, 6},
				NestedIntPointerList: &[]*[]int64{
					{int64(3), int64(4)},
				},
				NestedIntList: &[][]int64{
					{int64(3), int64(4)},
				},
				WindowedInt: map[string]*int64{
					"1m": ptr.New(int64(4)),
					"5m": ptr.New(int64(5)),
					"1h": ptr.New(int64(6)),
				},
				WindowedList: map[string]*[]int64{
					"1m": {4, 5, 6},
					"5m": {7, 8, 9},
					"1h": {10, 11, 12},
				},
				Dataclass: &fixtures.LatLng{
					Lat: ptr.New(5.5),
					Lng: ptr.New(6.6),
				},
				DataclassList: &[]fixtures.LatLng{
					{
						Lat: ptr.New(7.7),
						Lng: ptr.New(8.8),
					},
				},
				DataclassWithList: &fixtures.FavoriteThings{
					Numbers: &[]int64{4, 5, 6},
					Words:   &[]string{"d", "e", "f"},
				},
				DataclassWithNils: &fixtures.Possessions{
					Car:   ptr.New("car2"),
					Yacht: ptr.New("yacht2"),
					Plane: ptr.New("plane2"),
				},
				// CHA-5430
				//DataclassWithDataclass: &fixtures.Child{
				//	Name: ptr.New("child2"),
				//	Mom: &fixtures.Parent{
				//		Name: ptr.New("mom2"),
				//	},
				//},
				DataclassWithOverrides: &fixtures.DclassWithOverrides{
					CamelName: ptr.New("camel2"),
				},
			},
		})
	resolved, err := params.underlying.resolveSingle()
	assert.NoError(t, err)
	request, err := serializeOnlineQueryParams(&params.underlying, resolved)
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
	fixtureRoot, initErr := GetRootFeatures()
	assert.NoError(t, initErr)
	inputs := map[any]any{
		fixtureRoot.AllTypes.String: 1,
		fixtureRoot.AllTypes.Float:  1.1,
	}
	params := OnlineQueryParams{}.WithInputs(inputs)
	resolved, err := params.underlying.resolveSingle()
	assert.NoError(t, err)

	feature1, err := UnwrapFeature(fixtureRoot.AllTypes.String)
	assert.Nil(t, err)
	feature2, err := UnwrapFeature(fixtureRoot.AllTypes.Float)
	assert.Nil(t, err)

	_, ok := resolved.inputs[feature1.Fqn]
	assert.True(t, ok)
	_, ok = resolved.inputs[feature2.Fqn]
	assert.True(t, ok)
}

func TestWithInputsMapFromOnlineQueryParamsWithInputs(t *testing.T) {
	t.Parallel()
	fixtureRoot, initErr := GetRootFeatures()
	assert.NoError(t, initErr)
	inputs := map[any]any{
		fixtureRoot.AllTypes.String: 1,
		fixtureRoot.AllTypes.Float:  1.1,
	}
	params := OnlineQueryParams{}.WithInput(fixtureRoot.AllTypes.Bool, true).WithInputs(inputs)
	resolved, err := params.underlying.resolveSingle()
	assert.NoError(t, err)

	feature1, err := UnwrapFeature(fixtureRoot.AllTypes.String)
	assert.Nil(t, err)
	feature2, err := UnwrapFeature(fixtureRoot.AllTypes.Float)
	assert.Nil(t, err)

	_, ok := resolved.inputs[feature1.Fqn]
	assert.True(t, ok)
	_, ok = resolved.inputs[feature2.Fqn]
	assert.True(t, ok)
}

func TestWithInputsMapFromOnlineQueryParamsWithOutputs(t *testing.T) {
	t.Parallel()
	fixtureRoot, initErr := GetRootFeatures()
	assert.NoError(t, initErr)
	inputs := map[any]any{
		fixtureRoot.AllTypes.String: 1,
		fixtureRoot.AllTypes.Float:  1.1,
	}
	params := OnlineQueryParams{}.WithOutputs(fixtureRoot.AllTypes.Bool).WithInputs(inputs)

	resolved, err := params.underlying.resolveSingle()
	assert.NoError(t, err)

	feature1, err := UnwrapFeature(fixtureRoot.AllTypes.String)
	assert.Nil(t, err)
	feature2, err := UnwrapFeature(fixtureRoot.AllTypes.Float)
	assert.Nil(t, err)

	_, ok := resolved.inputs[feature1.Fqn]
	assert.True(t, ok)
	_, ok = resolved.inputs[feature2.Fqn]
	assert.True(t, ok)
}

func TestWithInputsMapFromOnlineQueryParamsComplete(t *testing.T) {
	t.Parallel()
	fixtureRoot, initErr := GetRootFeatures()
	assert.NoError(t, initErr)
	inputs := map[any]any{
		fixtureRoot.AllTypes.String: 1,
		fixtureRoot.AllTypes.Float:  1.1,
	}
	params := OnlineQueryParamsComplete{}.WithInputs(inputs)

	resolved, err := params.underlying.resolveSingle()
	assert.NoError(t, err)

	feature1, err := UnwrapFeature(fixtureRoot.AllTypes.String)
	assert.Nil(t, err)
	feature2, err := UnwrapFeature(fixtureRoot.AllTypes.Float)
	assert.Nil(t, err)

	_, ok := resolved.inputs[feature1.Fqn]
	assert.True(t, ok)
	_, ok = resolved.inputs[feature2.Fqn]
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

func TestInputsLengthValidation(t *testing.T) {
	t.Parallel()

	paramsSameLength := OnlineQueryParams{}.
		WithInput("user.id", []int{1, 2, 3}).
		WithInput("user.number", []int{1, 2, 3})

	_, err := paramsSameLength.underlying.resolveSingle()
	assert.NoError(t, err)
	_, err = paramsSameLength.underlying.resolveBulk()
	assert.NoError(t, err)

	paramsDiffLength := OnlineQueryParams{}.
		WithInput("user.id", []int{1, 2, 3}).
		WithInput("user.number", []int{1, 2})
	_, err = paramsDiffLength.underlying.resolveSingle()
	assert.NoError(t, err)
	_, err = paramsDiffLength.underlying.resolveBulk()
	assert.Error(t, err)
}
