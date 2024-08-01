package integration

import (
	"fmt"
	"github.com/chalk-ai/chalk-go"
	"github.com/samber/lo"
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
	Id                                  *string
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

// TestGrpcUploadFeatures is a separate test from the non-grpc upload test
// because the GRPC endpoint only handles uploading windowed agg aggregations.
func TestGrpcUploadFeatures(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)
	assert.NoError(t, chalk.InitFeatures(&Features))
	client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: true})
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}

	distinctProofIds := getRandomInts(4)
	proofIdBytes := [][]byte{
		[]byte(fmt.Sprintf("%d", distinctProofIds[0])),
		[]byte(fmt.Sprintf("%d", distinctProofIds[0])),
		[]byte(fmt.Sprintf("%d", distinctProofIds[1])),
		[]byte(fmt.Sprintf("%d", distinctProofIds[2])),
		[]byte(fmt.Sprintf("%d", distinctProofIds[3])),
	}
	theoremIds := getRandomInts(len(proofIdBytes))
	theoremIdsFloat := lo.Map(theoremIds, func(i int64, _ int) float64 { return float64(i) })

	now := time.Now().UTC()

	params := chalk.UploadFeaturesParams{Inputs: map[any]any{
		// This should be an int64 once the server automatically casts group by columns to float64
		Features.Theorem.Id: theoremIdsFloat,
		// This should be an int64 once the server automatically casts group by columns to binary
		Features.Theorem.ProofId: proofIdBytes,
		// This should be its original int64 type once the server automatically casts aggregated columns to float64
		Features.Theorem.NumLines: []float64{34, 55, 89, 144, 233},
		// This should be a string once the server automatically casts group by columns to binary
		Features.Theorem.Author: [][]byte{
			[]byte("Carl Friederich Gauss"),
			[]byte("Sir Isaac Newton"),
			[]byte("Albert Einstein"),
			[]byte("Richard Feynman"),
			[]byte("Sir Roger Penrose"),
		},
		"__ts__": []time.Time{
			now.Add(-1 * 24 * time.Hour),
			now.Add(-2 * 24 * time.Hour),
			now.Add(-3 * 24 * time.Hour),
			now.Add(-4 * 24 * time.Hour),
			now.Add(-5 * 24 * time.Hour),
		},
	}}

	_, err = client.UploadFeatures(params)
	assert.NoError(t, err)

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
