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
			t.Error("Standard deviation found", standardDeviation, ", expected close to 1")
		}
	})

	t.Run("Must match math/rand/Normal mean of 0", func(t *testing.T) {
		lowerTargetBound := -epsilon
		upperTargetBound := epsilon
		mean := CalculateMean(samples)

		if mean < lowerTargetBound || mean > upperTargetBound {
			t.Error("Mean found", mean, ", expected close to 0")
		}
	})
}
