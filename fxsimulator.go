package main

import (
	"flag"
	"fmt"

	"github.com/betocastelo/fxsimulator/fxsimlib"
	"os"
)

var (
	expectedForwardRateArg = flag.Float64("e", 0.0, "expected forward rate")
	volatilityArg          = flag.Float64("v", 0.0, "market volatility")
	timeToMaturityArg      = flag.Uint("t", 1, "time to maturity")
	lengthOfSimulationArg  = flag.Uint("l", 1, "number of runs of simulation")
)

func main() {
	flag.Parse()

	if (!inputIsValid()) {
		printUsageAndExit(1)
	}

	results := fxsim.FxForward(*expectedForwardRateArg, *volatilityArg,
		*timeToMaturityArg, *lengthOfSimulationArg)

	for _, result := range results {
		fmt.Println(result)
	}
}

func inputIsValid() bool {
	if len(flag.Args()) == 0 {
		return false
	}

	return true
}

func printUsageAndExit(code int) {
	flag.Usage()
	os.Exit(code)
}
