package uddsketch

import (
	"fmt"
	"math"
)

// SketchHashKey represents the bucket index in the UDDSketch
type SketchHashKey struct {
	keyType byte  // 0 = Invalid, 1 = Negative, 2 = Zero, 3 = Positive
	value   int64 // Used for Positive/Negative values
}

// EstimateQuantileAtValue estimates what quantile a value would be in the current sketch
func (s *UDDSketch) EstimateQuantileAtValue(value float64) float64 {
	if s.numValues == 0 {
		return 0.0
	}

	var count float64
	target := s.key(value)

	current := s.buckets.head
	for current.keyType != Invalid {
		entry := s.buckets.buckets[current]
		if lessThan(current, target) {
			count += float64(entry.count)
		} else if current == target {
			// If the value falls in the target bucket, assume it's greater than half the other values
			count += float64(entry.count) / 2.0
		}
		current = entry.next
	}

	return count / float64(s.numValues)
}

// MergeSketch merges another sketch into this one
func (s *UDDSketch) MergeSketch(other *UDDSketch) {
	// Require matching initial parameters
	gamma1 := math.Pow(s.gamma, 1.0/math.Pow(2.0, float64(s.compactions)))
	gamma2 := math.Pow(other.gamma, 1.0/math.Pow(2.0, float64(other.compactions)))
	if math.Abs(gamma1-gamma2) >= 1e-9 {
		return // silently fail like the Rust version
	}
	if s.maxBuckets != other.maxBuckets {
		return
	}

	// Handle empty sketches
	if other.numValues == 0 {
		return
	}
	if s.numValues == 0 {
		*s = *other
		return
	}

	// Create a copy of other sketch to manipulate
	otherCopy := *other

	// Align compaction levels
	for s.compactions > otherCopy.compactions {
		otherCopy.compactBuckets()
	}
	for otherCopy.compactions > s.compactions {
		s.compactBuckets()
	}

	// Merge buckets
	current := otherCopy.buckets.head
	for current.keyType != Invalid {
		entry := otherCopy.buckets.buckets[current]
		if existing, exists := s.buckets.buckets[current]; exists {
			existing.count += entry.count
			s.buckets.buckets[current] = existing
		} else {
			s.increment(current)
			s.buckets.buckets[current] = SketchHashEntry{
				count: entry.count,
				next:  s.buckets.buckets[current].next,
			}
		}
		current = entry.next
	}

	// Compact if needed
	for uint64(len(s.buckets.buckets)) > s.maxBuckets {
		s.compactBuckets()
	}

	// Update statistics
	s.numValues += other.numValues
	s.valuesSum += other.valuesSum
	s.min = math.Min(s.min, other.min)
	s.max = math.Max(s.max, other.max)
	s.zeroCount += other.zeroCount
}

// Constants for SketchHashKey types
const (
	Invalid byte = iota
	Negative
	Zero
	Positive
)

// SketchHashEntry represents an entry in the SketchHashMap
type SketchHashEntry struct {
	count uint64
	next  SketchHashKey
}

// SketchHashMap is a specialized hash map that maintains ordered buckets
type SketchHashMap struct {
	buckets map[SketchHashKey]SketchHashEntry
	head    SketchHashKey
}

// UDDSketch implements the UDDSketch algorithm
type UDDSketch struct {
	buckets     SketchHashMap
	alpha       float64
	gamma       float64
	compactions uint32
	maxBuckets  uint64
	numValues   uint64
	valuesSum   float64
	min         float64
	max         float64
	zeroCount   uint64
}

// NewUDDSketch creates a new UDDSketch with specified parameters
func NewUDDSketch(maxBuckets uint64, initialError float64) (*UDDSketch, error) {
	if initialError <= 0 || initialError >= 1.0 {
		return nil, fmt.Errorf("initial error must be between 0 and 1")
	}

	return &UDDSketch{
		buckets: SketchHashMap{
			buckets: make(map[SketchHashKey]SketchHashEntry),
			head:    SketchHashKey{keyType: Invalid},
		},
		alpha:      initialError,
		gamma:      (1.0 + initialError) / (1.0 - initialError),
		maxBuckets: maxBuckets,
		min:        math.MaxFloat64,
		max:        -math.MaxFloat64,
	}, nil
}

// key returns the appropriate bucket key for a given value
func (s *UDDSketch) key(value float64) SketchHashKey {
	if value == 0 {
		return SketchHashKey{keyType: Zero}
	}

	negative := value < 0
	value = math.Abs(value)
	logVal := math.Ceil(math.Log(value) / math.Log(s.gamma))

	if negative {
		return SketchHashKey{keyType: Negative, value: int64(logVal)}
	}
	return SketchHashKey{keyType: Positive, value: int64(logVal)}
}

// AddValue adds a value to the sketch
func (s *UDDSketch) AddValue(value float64) {
	if math.IsNaN(value) {
		return
	}

	key := s.key(value)
	s.increment(key)

	for uint64(len(s.buckets.buckets)) > s.maxBuckets {
		s.compactBuckets()
	}

	s.numValues++
	s.valuesSum += value
	s.min = math.Min(s.min, value)
	s.max = math.Max(s.max, value)
	if value == 0 {
		s.zeroCount++
	}
}

// increment increases the count for a given bucket
func (s *UDDSketch) increment(key SketchHashKey) {
	if entry, exists := s.buckets.buckets[key]; exists {
		// Just increment existing bucket
		entry.count++
		s.buckets.buckets[key] = entry
		return
	}

	// Handle new bucket
	var entry SketchHashEntry
	entry.count = 1

	if s.buckets.head.keyType == Invalid || lessThan(key, s.buckets.head) {
		// Insert at head
		entry.next = s.buckets.head
		s.buckets.head = key
	} else {
		// Find insertion point
		prev := s.buckets.head
		curr := s.buckets.buckets[prev].next
		for curr.keyType != Invalid && !lessThan(key, curr) {
			prev = curr
			curr = s.buckets.buckets[curr].next
		}
		entry.next = curr

		// Update previous bucket's next pointer
		prevEntry := s.buckets.buckets[prev]
		prevEntry.next = key
		s.buckets.buckets[prev] = prevEntry
	}
	s.buckets.buckets[key] = entry
}

// compactBuckets combines adjacent buckets to reduce storage
func (s *UDDSketch) compactBuckets() {
	newBuckets := make(map[SketchHashKey]SketchHashEntry)
	var newHead SketchHashKey
	newHead.keyType = Invalid

	current := s.buckets.head
	for current.keyType != Invalid {
		entry := s.buckets.buckets[current]
		newKey := s.compactKey(current)

		if newHead.keyType == Invalid {
			newHead = newKey
		}

		newEntry, exists := newBuckets[newKey]
		if !exists {
			newEntry = SketchHashEntry{count: 0}
		}
		newEntry.count += entry.count

		// Set next pointer
		nextKey := s.compactKey(entry.next)
		if nextKey == newKey {
			nextKey = s.compactKey(s.buckets.buckets[entry.next].next)
		}
		newEntry.next = nextKey
		newBuckets[newKey] = newEntry

		current = entry.next
	}

	s.buckets.buckets = newBuckets
	s.buckets.head = newHead
	s.compactions++
	s.gamma *= s.gamma
	s.alpha = 2.0 * s.alpha / (1.0 + s.alpha*s.alpha)
}

// compactKey returns the new key after compaction
func (s *UDDSketch) compactKey(key SketchHashKey) SketchHashKey {
	if key.keyType == Zero || key.keyType == Invalid {
		return key
	}

	value := key.value
	if value == math.MaxInt64 {
		return key
	}

	if value > 0 {
		value = (value + 1) / 2
	} else {
		value = value / 2
	}

	return SketchHashKey{keyType: key.keyType, value: value}
}

// EstimateQuantile returns the estimated value at the given quantile
func (s *UDDSketch) EstimateQuantile(quantile float64) float64 {
	if quantile < 0 || quantile > 1 {
		return math.NaN()
	}

	remaining := uint64(float64(s.numValues)*quantile) + 1
	if remaining >= s.numValues {
		return s.max
	}

	current := s.buckets.head
	for current.keyType != Invalid {
		entry := s.buckets.buckets[current]
		if remaining <= entry.count {
			return s.bucketToValue(current)
		}
		remaining -= entry.count
		current = entry.next
	}

	return math.NaN()
}

// bucketToValue converts a bucket key back to its representative value
func (s *UDDSketch) bucketToValue(key SketchHashKey) float64 {
	switch key.keyType {
	case Zero:
		return 0
	case Positive:
		return math.Pow(s.gamma, float64(key.value-1)) * (1.0 + s.alpha)
	case Negative:
		return -math.Pow(s.gamma, float64(key.value-1)) * (1.0 + s.alpha)
	default:
		return math.NaN()
	}
}

// Helper functions

func lessThan(a, b SketchHashKey) bool {
	if a.keyType != b.keyType {
		return a.keyType < b.keyType
	}
	if a.keyType == Positive {
		return a.value < b.value
	}
	if a.keyType == Negative {
		return a.value > b.value
	}
	return false
}

// Accessor methods

func (s *UDDSketch) Count() uint64 {
	return s.numValues
}

func (s *UDDSketch) Mean() float64 {
	if s.numValues == 0 {
		return 0
	}
	return s.valuesSum / float64(s.numValues)
}

func (s *UDDSketch) Sum() float64 {
	return s.valuesSum
}

func (s *UDDSketch) MaxError() float64 {
	return s.alpha
}

func (s *UDDSketch) Min() float64 {
	return s.min
}

func (s *UDDSketch) Max() float64 {
	return s.max
}

func (s *UDDSketch) ZeroCount() uint64 {
	return s.zeroCount
}
