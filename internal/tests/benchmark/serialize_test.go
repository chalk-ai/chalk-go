package benchmark

import (
	"fmt"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/chalk-ai/chalk-go/internal"
)

func serializeInputs(rows int) map[string]any {
	ids := make([]string, rows)
	ints := make([]int64, rows)
	fl := make([]float64, rows)
	ts := make([]time.Time, rows)
	for i := range ids {
		ids[i] = fmt.Sprintf("offer-%08d", i*7919)
		ints[i] = int64(i)
		fl[i] = float64(i) / 3
		ts[i] = time.Unix(1700000000+int64(i), 0).UTC()
	}
	return map[string]any{"user.id": ids, "user.n": ints, "user.f": fl, "user.t": ts}
}

func BenchmarkSerializeIdsOnly(b *testing.B) {
	in := map[string]any{"u.id": serializeInputs(1000)["user.id"]}
	b.ReportAllocs()
	for b.Loop() {
		if _, err := internal.InputsToArrowBytes(in, memory.DefaultAllocator); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSerializeMixed1000(b *testing.B) {
	in := serializeInputs(1000)
	b.ReportAllocs()
	for b.Loop() {
		if _, err := internal.InputsToArrowBytes(in, memory.DefaultAllocator); err != nil {
			b.Fatal(err)
		}
	}
}
