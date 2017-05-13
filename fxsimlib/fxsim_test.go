package fxsim

import (
	"math"
	"testing"

	"github.com/betocastelo/fxsimulator/simplestatslib"
)

type simParameters struct {
	expectedFwdRate    float64
	volatility         float64
	timeToMaturity     uint
	lengthOfSimulation uint
}

var (
	epsilon             = 0.005
	standardCase        = simParameters{1.5, .05, 10, 1000}
	zeroValueInputsCase = simParameters{0, 0, 0, 0}
)

func TestFxForwardCorrectness(t *testing.T) {
	t.Run("Must have mean = ln(F0)/T-.5*vol^2", func(t *testing.T) {
		samples := produceModifiedSimDataForTesting(standardCase)
		mean := simplestats.CalculateMean(samples)
		target := math.Log(standardCase.expectedFwdRate)/float64(standardCase.timeToMaturity) -
			0.5*math.Pow(standardCase.volatility, 2)
		lowerBoundTarget := target - epsilon
		upperBoundTarget := target + epsilon

		if mean < lowerBoundTarget || mean > upperBoundTarget {
			t.Error("mean =", mean, "instead of 0, within range of", epsilon)
		}
	})

	t.Run("Must have standardDeviation = volatility", func(t *testing.T) {
		samples := produceModifiedSimDataForTesting(standardCase)
		standardDeviation := simplestats.CalculateStandardDeviation(samples)
		lowerBoundTarget := standardCase.volatility - epsilon
		upperBoundTarget := standardCase.volatility + epsilon

		if standardDeviation < lowerBoundTarget || standardDeviation > upperBoundTarget {
			t.Error("standardDeviation =", standardDeviation, "instead of",
				standardCase.volatility, "within range of", epsilon)
		}
	})

	t.Run("Zeros as input produce zeros as output", func(t *testing.T) {
		samples := produceModifiedSimDataForTesting(zeroValueInputsCase)
		lowerBoundTarget := -epsilon
		upperBoundTarget := epsilon

		t.Run("Must have mean = 0", func(t *testing.T) {
			mean := simplestats.CalculateMean(samples)

			if mean < lowerBoundTarget || mean > upperBoundTarget {
				t.Error("mean =", mean, "instead of 0, within range of", epsilon)
			}
		})

		t.Run("Must have standardDeviation = 0", func(t *testing.T) {
			standardDeviation := simplestats.CalculateStandardDeviation(samples)

			if standardDeviation < lowerBoundTarget || standardDeviation > upperBoundTarget {
				t.Error("standardDeviation =", standardDeviation, "instead of 0, within range of",
					epsilon)
			}
		})
	})
}

func produceModifiedSimDataForTesting(simParms simParameters) []float64 {
	samples := FxForward(simParms.expectedFwdRate, simParms.volatility,
		simParms.timeToMaturity, simParms.lengthOfSimulation)
	modifyFxForward(samples, standardCase.timeToMaturity)
	return samples
}

// modifyFxForward takes each sample s of the output of FxForward and applies ln(s)/T, where
// T is the time-to-maturity parameter.
func modifyFxForward(forwardSim []float64, timeToMaturity uint) {
	for i, sample := range forwardSim {
		if sample != 0.0 {
			forwardSim[i] = math.Log(sample) / float64(timeToMaturity)
		}
	}
}
