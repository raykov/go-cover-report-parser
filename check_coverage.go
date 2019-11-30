package coverreportparser

import (
	"fmt"
)

func check(cov coverage, opts Options) (err error) {
	if opts.Verbose {
		cov.Print()
	}
	minCov := opts.MinimumExpectedCoverage

	if minCov > 100.0 || minCov < 0.0 {
		minCov = 100.0
	}

	if cov.Coverage() < minCov {
		fmt.Printf(
			"Coverage (%.2f%%) is below the expected minimum coverage (%.2f%%).",
			cov.Coverage(),
			minCov,
		)

		return UncoveredError
	} else {
		fmt.Printf("%.2f%% coverage\n", cov.Coverage())
	}

	return nil
}
