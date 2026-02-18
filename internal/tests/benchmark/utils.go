package benchmark

import (
	"sync"
	"testing"
	"time"
)

func benchmarkWithClose(b *testing.B, benchmarkFunc func(), closeFunc func()) {
	b.Helper()
	defer closeFunc()
	benchmark(b, benchmarkFunc)
	b.StopTimer()
}

func benchmarkParallelWithClose(b *testing.B, benchmarkFunc func(), closeFunc func()) {
	b.Helper()
	defer closeFunc()
	benchmarkParallel(b, benchmarkFunc)
	b.StopTimer()
}

func benchmark(b *testing.B, benchmarkFunc func()) {
	b.Helper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkFunc()
	}
	b.StopTimer()

	avg := b.Elapsed() / time.Duration(b.N)
	b.ReportMetric(0, "ns/op")                                  // Effective hides the default ns/op metric
	b.ReportMetric((float64(avg.Nanoseconds()) / 1e6), "ms/op") // The same metric but in ms
}

func benchmarkParallel(b *testing.B, benchmarkFunc func()) {
	b.Helper()
	numParallel := 200
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for range numParallel {
			wg.Go(func() {
				benchmarkFunc()
			})
		}
		wg.Wait()
	}
	b.StopTimer()
	avg := b.Elapsed() / time.Duration(b.N)
	b.ReportMetric(0, "ns/op")                                  // Effective hides the default ns/op metric
	b.ReportMetric((float64(avg.Nanoseconds()) / 1e6), "ms/op") // The same metric but in ms
}
