package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestCalculateStandardDeviationCorrectnessWithGoRandNormal(t *testing.T) {
	numOfSamples := 100000
	samples := make([]float64, numOfSamples)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(samples); i++ {
		samples[i] = r.NormFloat64()
	}

	epsilon := 0.01

	t.Run("Must match math/rand/Normal std dev of 1", func(t *testing.T) {
		lowerTargetBound := 1 - epsilon
		upperTargetBound := 1 + epsilon
		standardDeviation := CalculateStandardDeviation(samples)

		if standardDeviation < lowerTargetBound || standardDeviation > upperTargetBound {
			t.Errorf("Standard deviation found: %v, expected 1 ± %v", standardDeviation, epsilon)
		}
	})

	t.Run("Must match math/rand/Normal mean of 0", func(t *testing.T) {
		lowerTargetBound := -epsilon
		upperTargetBound := epsilon
		mean := CalculateMean(samples)

		if mean < lowerTargetBound || mean > upperTargetBound {
			t.Error("Mean found: %v, expected 0 ± %v", mean, epsilon)
		}
	})
}
