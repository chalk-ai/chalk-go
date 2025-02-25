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
	// Explicit NaN check needed in Go (Rust handles this implicitly)
	// Rust max/min operations ignore NaN values
	// Go wil return Nan for max/min operations on NaN values
	// Following the TimescaleDB implementation, NaN values are silently ignored.
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
	// Create new map for compacted buckets
	newBuckets := make(map[SketchHashKey]SketchHashEntry)
	var newHead SketchHashKey
	newHead.keyType = Invalid

	// Iterate through buckets in order
	current := s.buckets.head
	for current.keyType != Invalid {
		entry := s.buckets.buckets[current]
		// Get the new key after compaction for current bucket
		newKey := s.compactKey(current)

		// Set the head of our compacted structure on first iteration
		if newHead.keyType == Invalid {
			newHead = newKey
		}

		// Get or create entry for this compacted key
		newEntry, exists := newBuckets[newKey]
		if !exists {
			newEntry = SketchHashEntry{count: 0}
		}
		// Add counts from the old bucket to potentially existing counts
		newEntry.count += entry.count

		// Calculate the next pointer after compaction
		nextKey := s.compactKey(entry.next)
		if nextKey == newKey {
			// If the next bucket would compact to the same key as current,
			// skip it and use its next bucket as our next
			nextKey = s.compactKey(s.buckets.buckets[entry.next].next)
		}
		newEntry.next = nextKey
		newBuckets[newKey] = newEntry

		current = entry.next
	}

	// Replace old buckets with compacted version
	s.buckets.buckets = newBuckets
	s.buckets.head = newHead
	s.compactions++

	// Update gamma and alpha according to the paper's equations
	s.gamma *= s.gamma                                // Equation 3: γ_i+1 = γ_i^2
	s.alpha = 2.0 * s.alpha / (1.0 + s.alpha*s.alpha) // Equation 4: α_i+1 = 2α_i/(1 + α_i^2)
}

// compactKey returns the new key after compaction
func (s *UDDSketch) compactKey(key SketchHashKey) SketchHashKey {
	//Return Zero/Invalid keys unchanged
	//Return MaxInt64 keys unchanged
	//For positive values: (value + 1) / 2
	//For negative values: value / 2
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
		// Use cached max value instead of traversing to last bucket and computing its value
		// This is equivalent to Rust's last_bucket_value() but more efficient
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
	// Invalid is treated as greater than valid values
	if a.keyType == Invalid && b.keyType == Invalid {
		return false // Equal, so not less than
	}
	if a.keyType == Invalid {
		return false // Invalid is greater
	}
	if b.keyType == Invalid {
		return true // Invalid is greater
	}

	// Zero comparisons
	if a.keyType == Zero && b.keyType == Zero {
		return false // Equal, so not less than
	}

	// Positive comparisons
	if a.keyType == Positive && b.keyType == Positive {
		return a.value < b.value
	}

	// Negative comparisons
	if a.keyType == Negative && b.keyType == Negative {
		return a.value > b.value // Note: reversed comparison
	}

	// Cross-type comparisons
	if b.keyType == Positive {
		return true // Everything is less than Positive
	}
	if a.keyType == Positive {
		return false // Positive is greater than everything
	}
	if b.keyType == Negative {
		return false // Everything is greater than Negative
	}
	if a.keyType == Negative {
		return true // Negative is less than everything
	}

	return false // Shouldn't reach here
}

func (s *UDDSketch) ToRON() string {
	// There's no go-native RON format, so we'll just have to construct it manually
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
	return fmt.Sprintf("(version:%d,alpha:%s,max_buckets:%d,num_buckets:%d,compactions:%d,count:%d,sum:%s,buckets:[%s])",
		readable.Version,
		formatFloat(readable.Alpha),
		readable.MaxBuckets,
		readable.NumBuckets,
		readable.Compactions,
		readable.Count,
		formatFloat(readable.Sum),
		bucketStr,
	)
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
		} else {
			if current == target {
				// If the value falls in the target bucket, assume it's greater than half the other values
				count += float64(entry.count) / 2.0
			}
			return count / float64(s.numValues)
		}
		current = entry.next
	}

	return 1.0 // Value is greater than anything in the sketch
}

// MergeSketch merges another sketch into this one
func (s *UDDSketch) MergeSketch(other *UDDSketch) {
	// Require matching initial parameters
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

	// Create a deep copy of other sketch to manipulate
	otherCopy := *other
	otherCopy.buckets.buckets = make(map[SketchHashKey]SketchHashEntry, len(other.buckets.buckets))
	for k, v := range other.buckets.buckets {
		otherCopy.buckets.buckets[k] = v
	}

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
