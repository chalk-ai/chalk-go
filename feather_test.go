package chalk

import (
	"encoding/base64"
	"github.com/apache/arrow/go/v12/arrow"
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

func TestFeatherScalarsDeserialization(t *testing.T) {
	stringData, err := os.ReadFile("./internal/feather_scalars_test.txt")
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

	socureColName := (*result.ScalarData).ColumnName(2)
	assert.Equal(t, "user.socure_score", socureColName)
	socureCol := (*result.ScalarData).Column(2)
	assert.Equal(t, "[123 123 123]", socureCol.String())
	assert.Equal(t, arrow.PrimitiveTypes.Float64, socureCol.DataType())

	todayColName := (*result.ScalarData).ColumnName(3)
	assert.Equal(t, "user.today", todayColName)
	todayCol := (*result.ScalarData).Column(3)
	assert.Equal(t, "[1693180800000 1693180800000 1693180800000]", todayCol.String())
	assert.Equal(t, arrow.PrimitiveTypes.Float64, socureCol.DataType())
}

//gotest.mark.skip(reason="Fails at has-many parsing with EOF error. Why?")
//func TestFeatherGroupsDeserialization(t *testing.T) {
//	stringData, err := os.ReadFile("./internal/feather_groups_test.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	bytesData, err := base64.StdEncoding.DecodeString(string(stringData))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	bulkResponse := OnlineQueryBulkResponse{}
//	err = bulkResponse.Unmarshal(bytesData)
//	assert.Nil(t, err)
//
//	assert.Equal(t, 1, len(bulkResponse.QueryResults))
//	result := bulkResponse.QueryResults["0"]
//	assert.NotNil(t, result.ScalarData)
//
//}
