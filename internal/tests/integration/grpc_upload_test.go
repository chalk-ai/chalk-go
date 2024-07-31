package integration

import (
	"github.com/chalk-ai/chalk-go"
	assert "github.com/stretchr/testify/require"
	"testing"
	"time"
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
		Features.Theorem.Id: []int64{1, 2, 3, 4, 5},
		Features.Theorem.ProofId: [][]byte{
			[]byte("1"),
			[]byte("1"),
			[]byte("3"),
			[]byte("4"),
			[]byte("5"),
		},
		Features.Theorem.NumLines: []int64{34, 55, 89, 144, 233},
		Features.Theorem.Author: [][]byte{
			[]byte("Carl Friederich Gauss"),
			[]byte("Sir Isaac Newton"),
			[]byte("Albert Einstein"),
			[]byte("Richard Feynman"),
			[]byte("Sir Roger Penrose"),
		},
		"__ts__": []time.Time{
			time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 9, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 9, 3, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 9, 4, 0, 0, 0, 0, time.UTC),
			time.Date(2021, 9, 5, 0, 0, 0, 0, time.UTC),
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
