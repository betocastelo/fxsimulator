package simplestats

import (
	"math"
)

// CalculateMean does what is says (I have to write something here to keep go lint quiet)
func CalculateMean(samples []float64) float64 {
	total := 0.0
	for _, sample := range samples {
		total += sample
	}

	return total / float64(len(samples))
}

// CalculateStandardDeviation does what it says (I have to write something here to keep go lint
// quiet)
func CalculateStandardDeviation(samples []float64) float64 {
	mean := CalculateMean(samples)
	totalVariance := 0.0
	for _, sample := range samples {
		totalVariance += math.Pow(sample-mean, 2.0)
	}

	return math.Sqrt(totalVariance / float64(len(samples)))
}
