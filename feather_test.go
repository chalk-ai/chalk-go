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

	_, err := internal.CreateOnlineQueryBulkBody(params.underlying.inputs, params.underlying.outputs)
	assert.Nil(t, err)
}

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
}
