package main

import (
	"math"
	"math/rand"
	"time"
)

// FxForward produces an array of simulated forward rates given an expected forward rate f0, a
// volatility sigma, and a time-to-maturity T.
func FxForward(expectedForwardRate, volatility float64,
	timeToMaturity, lengthOfSimulation uint) []float64 {
	samples := make([]float64, lengthOfSimulation)

	// fixed components of the exponent
	a := volatility * float64(timeToMaturity)
	b := 0.5 * math.Pow(volatility, 2) * float64(timeToMaturity)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(samples); i++ {
		samples[i] = expectedForwardRate * math.Pow(math.E, -1.0*(r.NormFloat64()*a+b))
	}

	return samples
}
