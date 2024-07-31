package integration

import (
	"github.com/chalk-ai/assert"
	"github.com/chalk-ai/chalk-go"
	"testing"
)

type Theorem struct {
	Id       *int64
	ProofId  *string
	NumLines *int64
	Author   *string
}

/*
  count_theorem_lines: Windowed[int] = windowed(
        "30d",
        "90d",
        materialization={
            "bucket_duration": "1d",
        },
        expression=_.theorems.count(),
    )
    total_theorem_lines: Windowed[int] = windowed(
        "30d",
        "90d",
        materialization={
            "bucket_duration": "1d",
        },
        expression=_.theorems[_.num_lines].sum(),
    )
    mean_theorem_lines: Windowed[float] = windowed(
        "30d",
        "90d",
        materialization={
            "bucket_duration": "1d",
        },
        expression=_.theorems[_.num_lines].mean(),
    )

    num_theorem_lines_by_author: DataFrame = group_by_windowed(
        "30d",
        "90d",
        materialization={
            "bucket_duration": "1d",
        },
        expression=_.theorems.group_by(_.author).agg(_.num_lines.sum()),
    )

    num_theorem_lines_by_main_author_30d: int = feature(
        expression=_.num_theorem_lines_by_author["30d"].group(
            author=_.main_author,
        ),
    )
    num_theorem_lines_by_main_author_windowed: Windowed[int] = windowed(
        "30d",
        "90d",
        expression=_.num_theorem_lines_by_author.group(
            author=_.main_author,
        ),
    )
    main_author: str
*/

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

// TestGrpcUploadFeatures is a separate test from the non-grpc upload test
// because the GRPC endpoint only handles uploading windowed agg aggregations.
func TestGrpcUploadFeatures(t *testing.T) {
	t.Parallel()
	SkipIfNotIntegrationTester(t)

	assert.NoError(t, chalk.InitFeatures(&Features))
	// Implicitly sources config from env var
	client, err := chalk.NewClient(&chalk.ClientConfig{UseGrpc: true})
	if err != nil {
		t.Fatal("Failed creating a Chalk Client", err)
	}

	params := chalk.UploadFeaturesParams{Inputs: map[any]any{
		Features.Theorem.Id:       []int64{1, 2, 3, 4, 5},
		Features.Theorem.ProofId:  []string{"1", "1", "3", "4", "5"},
		Features.Theorem.NumLines: []int64{34, 55, 89, 144, 233},
		Features.Theorem.Author: []string{
			"Carl Friederich Gauss",
			"Sir Isaac Newton",
			"Albert Einstein",
			"Richard Feynman",
			"Sir Roger Penrose",
		},
		"__ts__": []string{
			"2021-09-01T00:00:00Z",
			"2021-09-02T00:00:00Z",
			"2021-09-03T00:00:00Z",
			"2021-09-04T00:00:00Z",
			"2021-09-05T00:00:00Z",
		},
	}}

	_, err = client.UploadFeatures(params)
	if err != nil {
		t.Fatal("Failed uploading features", err)
	}

	if err != nil {
		t.Fatal("Failed querying features", err)
	}

}
