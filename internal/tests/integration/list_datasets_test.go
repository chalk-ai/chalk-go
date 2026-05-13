package integration

import (
	"testing"

	chalk "github.com/chalk-ai/chalk-go"
	"github.com/stretchr/testify/assert"
)

func TestListDatasets(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	res, err := restClient.ListDatasets(t.Context(), chalk.ListDatasetsParams{Limit: 5})
	assert.NoError(t, err)
	assert.NotNil(t, res.RawResponse)

	res2, err := grpcClient.ListDatasets(t.Context(), chalk.ListDatasetsParams{Limit: 5})
	assert.NoError(t, err)
	assert.NotNil(t, res2.RawResponse)
}
