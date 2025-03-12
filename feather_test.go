package chalk

import (
	"encoding/base64"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestFeatherSerialization(t *testing.T) {
	params := OnlineQueryParams{}.
		WithInput("user.id", []int{1, 2, 3, 4}).
		WithOutputs("user.email", "user.card.id")
	resolved, err := params.underlying.resolveBulk()
	assert.NoError(t, err)
	_, err = internal.CreateOnlineQueryBulkBody(resolved.inputs, internal.FeatherRequestHeader{Outputs: resolved.outputs})
	assert.NoError(t, err)
}

// TestErrorDeserialization tests that we can successfully deserialize
// a feather response that contains errors.
func TestErrorDeserialization(t *testing.T) {
	stringData, err := os.ReadFile("./internal/sample_data/bulk_response_with_err.txt")
	if err != nil {
		log.Fatal(err)
	}

	bytesData, err := base64.StdEncoding.DecodeString(string(stringData))
	if err != nil {
		log.Fatal(err)
	}

	bulkResponse := OnlineQueryBulkResponse{}
	err = bulkResponse.Unmarshal(bytesData)
	assert.Nil(t, err)

	result := bulkResponse.QueryResults["0"]
	assert.NotNil(t, result.Errors)
	assert.Equal(t, 2, len(result.Errors))
	assert.Equal(t, result.Errors[0].Message, "Query abc referenced undefined feature 'def'")
	assert.Equal(t, result.Errors[1].Message, "query.ghi referenced invalid feature 'jkl'")
}

// TestFeatherDeserialization tests that we can successfully deserialize
// a feather response that contains no errors.
func TestFeatherDeserialization(t *testing.T) {
	// TODO: Add test for all data types
	stringData, err := os.ReadFile("./internal/sample_data/bulk_query_response.txt")
	if err != nil {
		log.Fatal(err)
	}

	bytesData, err := base64.StdEncoding.DecodeString(string(stringData))
	if err != nil {
		log.Fatal(err)
	}

	bulkResponse := OnlineQueryBulkResponse{}
	err = bulkResponse.Unmarshal(bytesData)
	assert.Nil(t, err)

	assert.Equal(t, 1, len(bulkResponse.QueryResults))
	result := bulkResponse.QueryResults["0"]
	assert.NotNil(t, result.ScalarData)
	assert.Equal(t, int64(5), result.ScalarData.NumRows())
	assert.NotNil(t, result.GroupsData)
	assert.NotNil(t, result.GroupsData["user.cups"])
	assert.Equal(t, int64(4), result.GroupsData["user.cups"].NumRows())
}
