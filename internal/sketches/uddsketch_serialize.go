package uddsketch

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
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
	// If it's a whole number with exactly .0, keep the .0
	if math.Mod(f, 1) == 0 {
		return fmt.Sprintf("%.1f", f)
	}

	// Convert to string and trim trailing zeros
	str := strconv.FormatFloat(f, 'f', -1, 64)
	return str
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

func decompressKeys(negativeIndexes []byte, hasZero bool, positiveIndexes []byte) []SketchHashKey {
	var keys []SketchHashKey
	decoder := newDeltaI64Encoder()

	offset := 0
	for offset < len(negativeIndexes) {
		deltaVal, n := decodeVarint(negativeIndexes[offset:])
		offset += n
		val := decoder.decode(deltaVal)
		keys = append(keys, SketchHashKey{keyType: Negative, value: val})
	}

	if hasZero {
		keys = append(keys, SketchHashKey{keyType: Zero})
	}

	decoder = newDeltaI64Encoder() // Reset for positive sequence
	offset = 0
	for offset < len(positiveIndexes) {
		deltaVal, n := decodeVarint(positiveIndexes[offset:])
		offset += n
		val := decoder.decode(deltaVal)
		keys = append(keys, SketchHashKey{keyType: Positive, value: val})
	}

	return keys
}

func decompressCounts(negativeCounts []byte, zeroBucketCount uint64, positiveCounts []byte) []uint64 {
	var counts []uint64
	decoder := newDeltaU64Encoder()

	offset := 0
	for offset < len(negativeCounts) {
		deltaVal, n := decodeVarint(negativeCounts[offset:])
		offset += n
		val := decoder.decode(uint64(deltaVal))
		counts = append(counts, val)
	}

	if zeroBucketCount > 0 {
		counts = append(counts, zeroBucketCount)
	}

	decoder = newDeltaU64Encoder() // Reset for positive sequence
	offset = 0
	for offset < len(positiveCounts) {
		deltaVal, n := decodeVarint(positiveCounts[offset:])
		offset += n
		val := decoder.decode(uint64(deltaVal))
		counts = append(counts, val)
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

type deltaI64Encoder struct {
	prev int64
}

type deltaU64Encoder struct {
	prev uint64
}

func newDeltaI64Encoder() *deltaI64Encoder {
	return &deltaI64Encoder{prev: 0}
}

func newDeltaU64Encoder() *deltaU64Encoder {
	return &deltaU64Encoder{prev: 0}
}

func (e *deltaI64Encoder) encode(val int64) int64 {
	delta := val - e.prev
	e.prev = val
	return delta
}

func (e *deltaI64Encoder) decode(delta int64) int64 {
	val := delta + e.prev
	e.prev = val
	return val
}

func (e *deltaU64Encoder) encode(val uint64) uint64 {
	delta := val - e.prev
	e.prev = val
	return delta
}

func (e *deltaU64Encoder) decode(delta uint64) uint64 {
	val := delta + e.prev
	e.prev = val
	return val
}

func compressBuckets(entries []BucketEntry) CompressedBuckets {
	var compressed CompressedBuckets
	indexEncoder := newDeltaI64Encoder()
	countEncoder := newDeltaU64Encoder()

	for _, entry := range entries {
		switch entry.Key.keyType {
		case Negative:
			deltaIdx := indexEncoder.encode(entry.Key.value)
			deltaCount := countEncoder.encode(entry.Count)
			compressed.negativeIndexes = append(compressed.negativeIndexes, encodeVarint(deltaIdx)...)
			compressed.negativeCounts = append(compressed.negativeCounts, encodeVarint(int64(deltaCount))...)
		case Zero:
			compressed.zeroBucketCount = entry.Count // No encoding needed
		case Positive:
			deltaIdx := indexEncoder.encode(entry.Key.value)
			deltaCount := countEncoder.encode(entry.Count)
			compressed.positiveIndexes = append(compressed.positiveIndexes, encodeVarint(deltaIdx)...)
			compressed.positiveCounts = append(compressed.positiveCounts, encodeVarint(int64(deltaCount))...)
		}
	}

	return compressed
}
