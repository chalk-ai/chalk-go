package integration

import (
	"context"
	"github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

type Theorem struct {
	Id       *int64
	ProofId  *int64
	NumLines *int64
	Author   *string
}

type Proof struct {
	Id                                  *int64
	CountTheoremLines                   map[string]*int64   `windows:"30d,90d"`
	TotalTheoremLines                   map[string]*int64   `windows:"30d,90d"`
	MeanTheoremLines                    map[string]*float64 `windows:"30d,90d"`
	NumTheoremLinesByAuthor             *map[string]*int64
	NumTheoremLinesByMainAuthor30d      *int64
	NumTheoremLinesByMainAuthorWindowed map[string]*int64 `windows:"30d,90d"`
	MainAuthor                          *string
}

var Features struct {
	Theorem *Theorem
	Proof   *Proof
}

func getRandomInts(n int) []int64 {
	var result []int64
	for i := 0; i < n; i++ {
		result = append(result, int64(rand.Int()))
	}
	return result
}

func verifyUpdateAggregateResults(t *testing.T, distinctProofIds []int64, client chalk.Client) {
	// Query aggregation results
	queryParams := chalk.OnlineQueryParams{}.WithInput(
		Features.Proof.Id,
		distinctProofIds,
	).WithOutputs(
		Features.Proof.TotalTheoremLines["30d"],
		Features.Proof.TotalTheoremLines["90d"],
		Features.Proof.MeanTheoremLines["30d"],
		Features.Proof.MeanTheoremLines["90d"],
	)
	bulkRes, err := client.OnlineQueryBulk(queryParams)
	assert.NoError(t, err)
	var proofResults []Proof
	err = bulkRes.UnmarshalInto(&proofResults)
	if err != (*chalk.ClientError)(nil) {
		t.Fatal("Failed querying features", err)
	}

	assert.Equal(t, 4, len(proofResults))
	assert.Equal(t, int64(89), *proofResults[0].TotalTheoremLines["30d"])
	assert.Equal(t, int64(89), *proofResults[0].TotalTheoremLines["90d"])
	assert.Equal(t, float64(44.5), *proofResults[0].MeanTheoremLines["30d"])
	assert.Equal(t, float64(44.5), *proofResults[0].MeanTheoremLines["90d"])

	assert.Equal(t, int64(89), *proofResults[1].TotalTheoremLines["30d"])
	assert.Equal(t, int64(89), *proofResults[1].TotalTheoremLines["90d"])
	assert.Equal(t, float64(89), *proofResults[1].MeanTheoremLines["30d"])
	assert.Equal(t, float64(89), *proofResults[1].MeanTheoremLines["90d"])

	assert.Equal(t, int64(144), *proofResults[2].TotalTheoremLines["30d"])
	assert.Equal(t, int64(144), *proofResults[2].TotalTheoremLines["90d"])
	assert.Equal(t, float64(144), *proofResults[2].MeanTheoremLines["30d"])
	assert.Equal(t, float64(144), *proofResults[2].MeanTheoremLines["90d"])

	assert.Equal(t, int64(233), *proofResults[3].TotalTheoremLines["30d"])
	assert.Equal(t, int64(233), *proofResults[3].TotalTheoremLines["90d"])
	assert.Equal(t, float64(233), *proofResults[3].MeanTheoremLines["30d"])
	assert.Equal(t, float64(233), *proofResults[3].MeanTheoremLines["90d"])
}

func getUpdateAggregateParams(proofIds []int64, theoremIds []int64, now time.Time) chalk.UpdateAggregatesParams {
	return chalk.UpdateAggregatesParams{Inputs: map[any]any{
		Features.Theorem.Id:       theoremIds,
		Features.Theorem.ProofId:  proofIds,
		Features.Theorem.NumLines: []int64{34, 55, 89, 144, 233},
		Features.Theorem.Author: []string{
			"Carl Friederich Gauss",
			"Sir Isaac Newton",
			"Albert Einstein",
			"Richard Feynman",
			"Sir Roger Penrose",
		},
		"theorem.__chalk_observed_at__": []time.Time{
			now.Add(-1 * 24 * time.Hour),
			now.Add(-2 * 24 * time.Hour),
			now.Add(-3 * 24 * time.Hour),
			now.Add(-4 * 24 * time.Hour),
			now.Add(-5 * 24 * time.Hour),
		},
	}}
}

// TestGrpcUpdateAggregates is a separate test from the non-grpc upload test
// because the GRPC endpoint only handles uploading windowed agg aggregations.
func TestGrpcUpdateAggregates(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	assert.NoError(t, chalk.InitFeatures(&Features))
	client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: true})
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}

	distinctProofIds := getRandomInts(4)
	proofIds := []int64{
		distinctProofIds[0],
		distinctProofIds[0],
		distinctProofIds[1],
		distinctProofIds[2],
		distinctProofIds[3],
	}
	theoremIds := getRandomInts(len(proofIds))
	now := time.Now().UTC()
	params := getUpdateAggregateParams(proofIds, theoremIds, now)
	_, err = client.UpdateAggregates(params)
	assert.NoError(t, err)
	// Query aggregation results
	verifyUpdateAggregateResults(t, distinctProofIds, client)
}

// TestGrpcUpdateAggregatesNative tests the same functionality as TestGrpcUpdateAggregates
// just from the native gRPC client.
func TestGrpcUpdateAggregatesNative(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	assert.NoError(t, chalk.InitFeatures(&Features))
	client, err := chalk.NewGRPCClient()
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}

	distinctProofIds := getRandomInts(4)
	proofIds := []int64{
		distinctProofIds[0],
		distinctProofIds[0],
		distinctProofIds[1],
		distinctProofIds[2],
		distinctProofIds[3],
	}
	theoremIds := getRandomInts(len(proofIds))
	now := time.Now().UTC()
	params := getUpdateAggregateParams(proofIds, theoremIds, now)
	_, err = client.UpdateAggregates(context.Background(), params)
	assert.NoError(t, err)

	restClient, err := chalk.NewClient()
	if err != nil {
		t.Fatal("Failed creating a REST client", err)
	}
	// Query aggregation results
	verifyUpdateAggregateResults(t, distinctProofIds, restClient)
}
