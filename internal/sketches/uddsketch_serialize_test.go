package uddsketch

import (
	"testing"
)

func TestRoundtripSerialize(t *testing.T) {
	sketch, err := NewUDDSketch(5, 0.000001)
	if err != nil {
		t.Fatalf("Failed to create sketch: %v", err)
	}

	for i := 0; i < 5; i++ {
		sketch.AddValue(float64(i))
	}

	serialized := sketch.ToRON()
	expected := "(version:1,alpha:0.000001,max_buckets:5,num_buckets:5,compactions:0,count:5,sum:10.0,buckets:[(Zero,1),(Positive(0),1),(Positive(346574),1),(Positive(549307),1),(Positive(693148),1)])"

	if serialized != expected {
		t.Errorf("Serialization mismatch:\nGot:      %s\nExpected: %s", serialized, expected)
	}

	// Also verify the fundamental properties
	readable := sketch.ToReadable()
	if sketch.Count() != readable.Count {
		t.Errorf("Count mismatch: got %d, want %d", readable.Count, sketch.Count())
	}
	if sketch.MaxError() != readable.Alpha {
		t.Errorf("Alpha mismatch: got %f, want %f", readable.Alpha, sketch.MaxError())
	}
	if uint32(len(sketch.buckets.buckets)) != readable.NumBuckets {
		t.Errorf("NumBuckets mismatch: got %d, want %d", readable.NumBuckets, len(sketch.buckets.buckets))
	}
	if uint64(sketch.compactions) != readable.Compactions {
		t.Errorf("Compactions mismatch: got %d, want %d", readable.Compactions, sketch.compactions)
	}
	if uint32(sketch.maxBuckets) != readable.MaxBuckets {
		t.Errorf("MaxBuckets mismatch: got %d, want %d", readable.MaxBuckets, sketch.maxBuckets)
	}
	if sketch.Sum() != readable.Sum {
		t.Errorf("Sum mismatch: got %f, want %f", readable.Sum, sketch.Sum())
	}
}
