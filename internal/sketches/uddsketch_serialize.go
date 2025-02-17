package uddsketch

import (
	"encoding/binary"
	"fmt"
	"math"
)

// PgUddSketch represents the PostgreSQL-compatible UDDSketch structure
type PgUddSketch struct {
	alpha           float64
	maxBuckets      uint32
	numBuckets      uint32
	compactions     uint64
	count           uint64
	sum             float64
	zeroBucketCount uint64
	negIndexesBytes uint32
	negBucketsBytes uint32
	posIndexesBytes uint32
	posBucketsBytes uint32
	negativeIndexes []byte
	negativeCounts  []byte
	positiveIndexes []byte
	positiveCounts  []byte
}

// ReadableUddSketch represents a human-readable format of the UDDSketch
type ReadableUddSketch struct {
	Version     uint8         `json:"version"`
	Alpha       float64       `json:"alpha"`
	MaxBuckets  uint32        `json:"max_buckets"`
	NumBuckets  uint32        `json:"num_buckets"`
	Compactions uint64        `json:"compactions"`
	Count       uint64        `json:"count"`
	Sum         float64       `json:"sum"`
	Buckets     []BucketEntry `json:"buckets"`
}

// BucketEntry represents a bucket in the sketch
type BucketEntry struct {
	Key   SketchHashKey `json:"key"`
	Count uint64        `json:"count"`
}

// CompressedBuckets holds the compressed representation of sketch buckets
type CompressedBuckets struct {
	negativeIndexes []byte
	negativeCounts  []byte
	zeroBucketCount uint64
	positiveIndexes []byte
	positiveCounts  []byte
}

func (s *UDDSketch) ToReadable() *ReadableUddSketch {
	pg := fromInternal(s)
	return fromPgSketch(pg)
}

func formatFloat(f float64) string {
	// Check if it's a whole number by comparing with its integer part
	if math.Floor(f) == f {
		return fmt.Sprintf("%.1f", f) // Will give X.0 for whole numbers
	}
	return fmt.Sprintf("%f", f)
}

func (s *UDDSketch) ToRON() string {
	readable := s.ToReadable()

	// Format buckets in RON style
	bucketStr := ""
	for i, b := range readable.Buckets {
		if i > 0 {
			bucketStr += ","
		}
		if b.Key.keyType == Zero {
			bucketStr += "(Zero," + fmt.Sprintf("%d", b.Count) + ")"
		} else if b.Key.keyType == Positive {
			bucketStr += "(Positive(" + fmt.Sprintf("%d", b.Key.value) + ")," + fmt.Sprintf("%d", b.Count) + ")"
		} else if b.Key.keyType == Negative {
			bucketStr += "(Negative(" + fmt.Sprintf("%d", b.Key.value) + ")," + fmt.Sprintf("%d", b.Count) + ")"
		}
	}

	// Format the entire string in RON style
	return fmt.Sprintf("(version:%d,alpha:%f,max_buckets:%d,num_buckets:%d,compactions:%d,count:%d,sum:%s,buckets:[%s])",
		readable.Version,
		readable.Alpha,
		readable.MaxBuckets,
		readable.NumBuckets,
		readable.Compactions,
		readable.Count,
		formatFloat(readable.Sum),
		bucketStr)
}

func fromInternal(state *UDDSketch) *PgUddSketch {
	var bucketEntries []BucketEntry
	current := state.buckets.head
	for current.keyType != Invalid {
		entry := state.buckets.buckets[current]
		bucketEntries = append(bucketEntries, BucketEntry{
			Key:   current,
			Count: entry.count,
		})
		current = entry.next
	}

	compressed := compressBuckets(bucketEntries)

	return &PgUddSketch{
		alpha:           state.MaxError(),
		maxBuckets:      uint32(state.maxBuckets),
		numBuckets:      uint32(len(state.buckets.buckets)),
		compactions:     uint64(state.compactions),
		count:           state.numValues,
		sum:             state.valuesSum,
		zeroBucketCount: compressed.zeroBucketCount,
		negIndexesBytes: uint32(len(compressed.negativeIndexes)),
		negBucketsBytes: uint32(len(compressed.negativeCounts)),
		posIndexesBytes: uint32(len(compressed.positiveIndexes)),
		posBucketsBytes: uint32(len(compressed.positiveCounts)),
		negativeIndexes: compressed.negativeIndexes,
		negativeCounts:  compressed.negativeCounts,
		positiveIndexes: compressed.positiveIndexes,
		positiveCounts:  compressed.positiveCounts,
	}
}

func fromPgSketch(pg *PgUddSketch) *ReadableUddSketch {
	var buckets []BucketEntry

	keys := decompressKeys(pg.negativeIndexes, pg.zeroBucketCount != 0, pg.positiveIndexes)
	counts := decompressCounts(pg.negativeCounts, pg.zeroBucketCount, pg.positiveCounts)

	for i := range keys {
		buckets = append(buckets, BucketEntry{
			Key:   keys[i],
			Count: counts[i],
		})
	}

	return &ReadableUddSketch{
		Version:     1,
		Alpha:       pg.alpha,
		MaxBuckets:  pg.maxBuckets,
		NumBuckets:  pg.numBuckets,
		Compactions: pg.compactions,
		Count:       pg.count,
		Sum:         pg.sum,
		Buckets:     buckets,
	}
}

func compressBuckets(entries []BucketEntry) CompressedBuckets {
	var compressed CompressedBuckets

	for _, entry := range entries {
		switch entry.Key.keyType {
		case Negative:
			compressed.negativeIndexes = append(compressed.negativeIndexes, encodeVarint(entry.Key.value)...)
			compressed.negativeCounts = append(compressed.negativeCounts, encodeVarint(int64(entry.Count))...)
		case Zero:
			compressed.zeroBucketCount = entry.Count
		case Positive:
			compressed.positiveIndexes = append(compressed.positiveIndexes, encodeVarint(entry.Key.value)...)
			compressed.positiveCounts = append(compressed.positiveCounts, encodeVarint(int64(entry.Count))...)
		}
	}

	return compressed
}

func decompressKeys(negativeIndexes []byte, hasZero bool, positiveIndexes []byte) []SketchHashKey {
	var keys []SketchHashKey

	offset := 0
	for offset < len(negativeIndexes) {
		val, n := decodeVarint(negativeIndexes[offset:])
		offset += n
		keys = append(keys, SketchHashKey{keyType: Negative, value: val})
	}

	if hasZero {
		keys = append(keys, SketchHashKey{keyType: Zero})
	}

	offset = 0
	for offset < len(positiveIndexes) {
		val, n := decodeVarint(positiveIndexes[offset:])
		offset += n
		keys = append(keys, SketchHashKey{keyType: Positive, value: val})
	}

	return keys
}

func decompressCounts(negativeCounts []byte, zeroBucketCount uint64, positiveCounts []byte) []uint64 {
	var counts []uint64

	offset := 0
	for offset < len(negativeCounts) {
		val, n := decodeVarint(negativeCounts[offset:])
		offset += n
		counts = append(counts, uint64(val))
	}

	if zeroBucketCount > 0 {
		counts = append(counts, zeroBucketCount)
	}

	offset = 0
	for offset < len(positiveCounts) {
		val, n := decodeVarint(positiveCounts[offset:])
		offset += n
		counts = append(counts, uint64(val))
	}

	return counts
}

func encodeVarint(x int64) []byte {
	var buf [10]byte
	n := binary.PutVarint(buf[:], x)
	return buf[:n]
}

func decodeVarint(buf []byte) (int64, int) {
	return binary.Varint(buf)
}
