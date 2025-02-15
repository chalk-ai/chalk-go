package uddsketch

import (
	"math"
	"math/rand"
	"testing"
)

const epsilon = 1e-3

func TestBuildAndAddValues(t *testing.T) {
	sketch, err := NewUDDSketch(20, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch: %v", err)
	}

	sketch.AddValue(1.0)
	sketch.AddValue(3.0)
	sketch.AddValue(0.5)

	if count := sketch.Count(); count != 3 {
		t.Errorf("Expected count 3, got %d", count)
	}
	if mean := sketch.Mean(); math.Abs(mean-1.5) > 1e-10 {
		t.Errorf("Expected mean 1.5, got %f", mean)
	}
	if maxError := sketch.MaxError(); math.Abs(maxError-0.1) > 1e-10 {
		t.Errorf("Expected max error 0.1, got %f", maxError)
	}
}

func TestExceedBuckets(t *testing.T) {
	sketch, err := NewUDDSketch(20, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch: %v", err)
	}

	sketch.AddValue(1.1)   // Bucket #1
	sketch.AddValue(400.0) // Bucket #30

	a2 := 0.2 / 1.01

	if count := sketch.Count(); count != 2 {
		t.Errorf("Expected count 2, got %d", count)
	}
	if maxError := sketch.MaxError(); math.Abs(maxError-0.1) > 1e-10 {
		t.Errorf("Expected max error 0.1, got %f", maxError)
	}

	for i := 2; i < 20; i++ {
		sketch.AddValue(1000.0 * math.Pow(1.23, float64(i)))
	}

	if count := sketch.Count(); count != 20 {
		t.Errorf("Expected count 20, got %d", count)
	}
	if maxError := sketch.MaxError(); math.Abs(maxError-0.1) > 1e-10 {
		t.Errorf("Expected max error 0.1, got %f", maxError)
	}

	for i := 20; i < 30; i++ {
		sketch.AddValue(1000.0 * math.Pow(1.23, float64(i)))
	}

	if count := sketch.Count(); count != 30 {
		t.Errorf("Expected count 30, got %d", count)
	}
	if maxError := sketch.MaxError(); math.Abs(maxError-a2) > 1e-10 {
		t.Errorf("Expected max error %f, got %f", a2, maxError)
	}
}

func TestMergeSketches(t *testing.T) {
	a1 := 0.1                                // alpha for up to 20 buckets
	a2 := 0.2 / 1.01                         // alpha for 1 compaction
	a3 := 2.0 * a2 / (1.0 + math.Pow(a2, 2)) // alpha for 2 compactions
	a4 := 2.0 * a3 / (1.0 + math.Pow(a3, 2)) // alpha for 3 compactions
	a5 := 2.0 * a4 / (1.0 + math.Pow(a4, 2)) // alpha for 4 compactions

	sketch1, err := NewUDDSketch(20, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch1: %v", err)
	}
	sketch1.AddValue(1.1) // Bucket #1
	sketch1.AddValue(1.5) // Bucket #3
	sketch1.AddValue(1.6) // Bucket #3
	sketch1.AddValue(1.3) // Bucket #2
	sketch1.AddValue(4.2) // Bucket #8

	if count := sketch1.Count(); count != 5 {
		t.Errorf("Expected count 5, got %d", count)
	}
	if maxError := sketch1.MaxError(); math.Abs(maxError-a1) > 1e-10 {
		t.Errorf("Expected max error %f, got %f", a1, maxError)
	}

	sketch2, err := NewUDDSketch(20, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch2: %v", err)
	}
	sketch2.AddValue(5.1)  // Bucket #9
	sketch2.AddValue(7.5)  // Bucket #11
	sketch2.AddValue(10.6) // Bucket #12
	sketch2.AddValue(9.3)  // Bucket #12
	sketch2.AddValue(11.2) // Bucket #13

	if maxError := sketch2.MaxError(); math.Abs(maxError-a1) > 1e-10 {
		t.Errorf("Expected max error %f, got %f", a1, maxError)
	}

	sketch1.MergeSketch(sketch2)
	if count := sketch1.Count(); count != 10 {
		t.Errorf("Expected count 10, got %d", count)
	}
	if maxError := sketch1.MaxError(); math.Abs(maxError-a1) > 1e-10 {
		t.Errorf("Expected max error %f, got %f", a1, maxError)
	}

	sketch3, err := NewUDDSketch(20, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch3: %v", err)
	}
	sketch3.AddValue(0.8)  // Bucket #-1
	sketch3.AddValue(3.7)  // Bucket #7
	sketch3.AddValue(15.2) // Bucket #14
	sketch3.AddValue(3.4)  // Bucket #7
	sketch3.AddValue(0.6)  // Bucket #-2

	if maxError := sketch3.MaxError(); math.Abs(maxError-a1) > 1e-10 {
		t.Errorf("Expected max error %f, got %f", a1, maxError)
	}

	sketch1.MergeSketch(sketch3)
	if count := sketch1.Count(); count != 15 {
		t.Errorf("Expected count 15, got %d", count)
	}
	if maxError := sketch1.MaxError(); math.Abs(maxError-a1) > 1e-10 {
		t.Errorf("Expected max error %f, got %f", a1, maxError)
	}

	sketch4, err := NewUDDSketch(20, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch4: %v", err)
	}
	sketch4.AddValue(400.0)           // Bucket #30
	sketch4.AddValue(0.004)           // Bucket #-27
	sketch4.AddValue(0.0)             // Zero Bucket
	sketch4.AddValue(-400.0)          // Neg. Bucket #30
	sketch4.AddValue(-0.004)          // Neg. Bucket #-27
	sketch4.AddValue(400000000000.0)  // Some arbitrary large bucket
	sketch4.AddValue(0.00000005)      // Some arbitrary small bucket
	sketch4.AddValue(-400000000000.0) // Some arbitrary large neg. bucket
	sketch4.AddValue(-0.00000005)     // Some arbitrary small neg. bucket

	if maxError := sketch4.MaxError(); math.Abs(maxError-a1) > 1e-10 {
		t.Errorf("Expected max error %f, got %f", a1, maxError)
	}

	sketch1.MergeSketch(sketch4)
	if count := sketch1.Count(); count != 24 {
		t.Errorf("Expected count 24, got %d", count)
	}
	if maxError := sketch1.MaxError(); math.Abs(maxError-a2) > 1e-10 {
		t.Errorf("Expected max error %f, got %f", a2, maxError)
	}

	sketch5, err := NewUDDSketch(20, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch5: %v", err)
	}
	for i := 100; i < 220; i++ {
		sketch5.AddValue(math.Pow(1.23, float64(i)))
	}

	if maxError := sketch5.MaxError(); math.Abs(maxError-a4) > 1e-10 {
		t.Errorf("Expected max error %f, got %f", a4, maxError)
	}

	sketch1.MergeSketch(sketch5)
	if count := sketch1.Count(); count != 144 {
		t.Errorf("Expected count 144, got %d", count)
	}
	if maxError := sketch1.MaxError(); math.Abs(maxError-a5) > 1e-10 {
		t.Errorf("Expected max error %f, got %f", a5, maxError)
	}
}

func TestZeroCount(t *testing.T) {
	sketch1, err := NewUDDSketch(50, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch1: %v", err)
	}

	for v := 0; v <= 10000; v++ {
		sketch1.AddValue(float64(v))
	}
	if count := sketch1.ZeroCount(); count != 1 {
		t.Errorf("Expected zero count 1, got %d", count)
	}

	sketch2, err := NewUDDSketch(50, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch2: %v", err)
	}
	for v := 0; v <= 10; v++ {
		sketch2.AddValue(float64(v) * math.SmallestNonzeroFloat64)
	}
	if count := sketch2.ZeroCount(); count != 1 {
		t.Errorf("Expected zero count 1, got %d", count)
	}

	sketch1.MergeSketch(sketch2)
	if count := sketch1.ZeroCount(); count != 2 {
		t.Errorf("Expected zero count 2, got %d", count)
	}
}

func TestQuantileAndValueEstimates(t *testing.T) {
	sketch, err := NewUDDSketch(50, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch: %v", err)
	}

	for v := 1; v <= 10000; v++ {
		sketch.AddValue(float64(v) / 100.0)
	}

	if count := sketch.Count(); count != 10000 {
		t.Errorf("Expected count 10000, got %d", count)
	}
	if maxError := sketch.MaxError(); math.Abs(maxError-0.1) > 1e-10 {
		t.Errorf("Expected max error 0.1, got %f", maxError)
	}

	for i := 1; i <= 100; i++ {
		value := float64(i)
		quantile := value / 100.0
		quantileValue := value + 0.01 // correct value for quantile should be next number > value

		testValue := sketch.EstimateQuantile(quantile)
		testQuant := sketch.EstimateQuantileAtValue(value)

		percentage := math.Abs(testValue-quantileValue) / quantileValue
		if percentage > 0.1 {
			t.Errorf("Exceeded 10%% error on quantile %f: expected %f, received %f (error%% %f)",
				quantile, quantileValue, testValue, percentage)
		}

		percentage = math.Abs(testQuant-quantile) / quantile
		if percentage > 0.2 {
			t.Errorf("Exceeded 20%% error on quantile at value %f: expected %f, received %f (error%% %f)",
				value, quantile, testQuant, percentage)
		}
	}

	if mean := sketch.Mean(); math.Abs(mean-50.005) >= epsilon {
		t.Errorf("Expected mean around 50.005, got %f", mean)
	}
	if minVal := sketch.Min(); math.Abs(minVal-0.01) >= epsilon {
		t.Errorf("Expected min around 0.01, got %f", minVal)
	}
	if maxVal := sketch.Max(); math.Abs(maxVal-100.0) >= epsilon {
		t.Errorf("Expected max around 100.0, got %f", maxVal)
	}
	if count := sketch.ZeroCount(); count != 0 {
		t.Errorf("Expected zero count 0, got %d", count)
	}
}

func TestExtremeQuantileAtValue(t *testing.T) {
	sketch, err := NewUDDSketch(50, 0.1)
	if err != nil {
		t.Fatalf("Failed to create sketch: %v", err)
	}

	for v := 1; v <= 10000; v++ {
		sketch.AddValue(float64(v) / 100.0)
	}

	extremeTests := map[float64]float64{
		-100.0: 0.0,
		0.0:    0.0,
		0.0001: 0.0,
		1000.0: 1.0,
	}

	for input, expected := range extremeTests {
		if result := sketch.EstimateQuantileAtValue(input); math.Abs(result-expected) > 1e-10 {
			t.Errorf("For input %f, expected %f, got %f", input, expected, result)
		}
	}

	// Additional range checks
	if result := sketch.EstimateQuantileAtValue(0.01); result >= 0.0001 {
		t.Errorf("Expected value < 0.0001 for input 0.01, got %f", result)
	}
	if result := sketch.EstimateQuantileAtValue(100.0); result <= 0.9 {
		t.Errorf("Expected value > 0.9 for input 100.0, got %f", result)
	}
}

func TestRandomStress(t *testing.T) {
	sketch, err := NewUDDSketch(1000, 0.01)
	if err != nil {
		t.Fatalf("Failed to create sketch: %v", err)
	}

	seed := int64(12345) // Using fixed seed from Rust test
	rng := rand.New(rand.NewSource(seed))
	bounds := make([]float64, 100)
	for i := range bounds {
		bounds[i] = rng.Float64()*2000000.0 - 1000000.0
		sketch.AddValue(bounds[i])
	}

	sortFloats(bounds)

	prev := -2000000.0
	for _, f := range bounds {
		for i := 0; i < 10000; i++ {
			sketch.AddValue(prev + rng.Float64()*(f-prev))
		}
		prev = f
	}

	for i := 0; i < 100; i++ {
		q := (float64(i) + 1.0) / 100.0
		val := sketch.EstimateQuantile(q)
		target := bounds[i]
		if target == 0 {
			continue
		}
		relError := math.Abs((val/target)-1.0) / math.Abs(target)
		if relError > sketch.MaxError() {
			t.Errorf("Failed to correctly match %f quantile with seed %d. Received: %f, Expected: %f, Error: %f, Expected error bound: %f",
				q, seed, val, target, relError, sketch.MaxError())
		}
	}
}

// Helper function to sort float64 slice
func sortFloats(a []float64) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
func TestFuzzing(t *testing.T) {
	rng := rand.New(rand.NewSource(12345))

	// Match the Rust test exactly - generate 4 batches of values
	generateBatch := func() []float64 {
		seen := make(map[float64]bool)
		for len(seen) < 25 {
			// Use exact same range as Rust test
			v := rng.Float64()*200.0 - 100.0
			if !math.IsNaN(v) {
				seen[v] = true
			}
		}
		result := make([]float64, 0, len(seen))
		for v := range seen {
			result = append(result, v)
		}
		return result
	}

	// Generate master list like Rust test
	var master []float64
	for i := 0; i < 4; i++ {
		master = append(master, generateBatch()...)
	}

	if len(master) < 100 {
		t.Skip("Not enough values for meaningful test")
	}

	sketch, err := NewUDDSketch(100, 0.000001)
	if err != nil {
		t.Fatalf("Failed to create sketch: %v", err)
	}

	for _, value := range master {
		sketch.AddValue(value)
	}

	sortFloats(master)

	// Test exactly the same quantiles as Rust
	quantileTests := []float64{0.01, 0.1, 0.25, 0.5, 0.6, 0.8, 0.95}
	for _, quantile := range quantileTests {
		// Compute target quantile using nearest rank definition
		masterIdx := int(math.Floor(quantile * float64(len(master))))
		target := master[masterIdx]

		if math.IsInf(target, 0) {
			continue
		}

		testVal := sketch.EstimateQuantile(quantile)
		// Handle infinite values like Rust test
		if math.IsInf(testVal, 0) {
			if testVal < 0 {
				testVal = math.MinInt64
			} else {
				testVal = math.MaxInt64
			}
		}

		var errVal float64
		if target == 0 {
			errVal = testVal
		} else {
			errVal = math.Abs(testVal-target) / math.Abs(target)
		}

		if errVal > sketch.MaxError() {
			t.Logf("master values: %v", master)
			t.Errorf("sketch with error %f estimated %f quantile as %f, true value is %f resulting in relative error %f",
				sketch.MaxError(), quantile, testVal, target, errVal)
		}
	}
}
