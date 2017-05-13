package main

import (
	"flag"
	"fmt"

	"github.com/betocastelo/fxsimulator/fxsimlib"
)

var (
	expectedForwardRateArg = flag.Float64("e", 0.0, "expected forward rate")
	volatilityArg          = flag.Float64("v", 0.0, "market volatility")
	timeToMaturityArg      = flag.Uint("t", 1, "time to maturity")
	lengthOfSimulationArg  = flag.Uint("l", 1, "number of runs of simulation")
)

func main() {
	flag.Parse()
	results := fxsim.FxForward(*expectedForwardRateArg, *volatilityArg,
		*timeToMaturityArg, *lengthOfSimulationArg)

	for _, result := range results {
		fmt.Println(result)
	}
}
