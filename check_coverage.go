package main

import (
	"fmt"
)

func checkCoverage(cov coverage) (err error) {
	if *verbose || *v {
		cov.Print()
	}

	if *minimumExpectedCoverage > 100.0 || *minimumExpectedCoverage < 0.0 {
		*minimumExpectedCoverage = 100.0
	}

	if cov.Coverage() < *minimumExpectedCoverage {
		fmt.Printf("Coverage (%.2f%%) is below the expected minimum coverage (%.2f%%).", cov.Coverage(), *minimumExpectedCoverage)
		return uncoveredError
	} else {
		fmt.Printf("%.2f%% coverage\n", cov.Coverage())
	}

	return
}
