package benchmark

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/tests/fixtures"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func getBenchmarkConvertBytesToTable(b *testing.B) func() {
	prefix := "really.long.prefix.of.feature"
	inputs := map[string]any{}

	numCols := 40
	for j := range numCols {
		inputs[fmt.Sprintf("%s%d", prefix, j)] = []int{}
	}

	numRows := 10_000

	for i := range numRows {
		for j := range numCols {
			inputs[fmt.Sprintf("%s%d", prefix, j)] = append(inputs[fmt.Sprintf("%s%d", prefix, j)].([]int), i)
		}
	}

	bytes, err := internal.InputsToArrowBytes(inputs, fixtures.TestAllocator)
	assert.NoError(b, err)

	return func() {
		_, err := internal.ConvertBytesToTable(bytes, fixtures.TestAllocator)
		assert.NoError(b, err)
	}
}

func BenchmarkConvertBytesToTable(b *testing.B) {
	benchmark(b, getBenchmarkConvertBytesToTable(b))
}

func BenchmarkConvertBytesToTableParallel(b *testing.B) {
	benchmarkParallel(b, getBenchmarkConvertBytesToTable(b))
}
