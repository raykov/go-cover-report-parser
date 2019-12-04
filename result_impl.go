package coverreportparser

import "fmt"

type result struct {
	minCoverage float64
	verbose     bool
	report      Report
}

func (r *result) String() string {
	str := ""

	if r.verbose {
		str += r.report.String()
	}

	coverage := r.report.Coverage()

	if coverage < r.minCoverage {
		str += fmt.Sprintf(
			"Coverage (%.2f%%) is below the expected minimum coverage (%.2f%%).\n",
			coverage,
			r.minCoverage,
		)
	} else {
		str += fmt.Sprintf("%.2f%% coverage\n", coverage)
	}

	return str
}

func (r *result) IsFailure() bool {
	return r.report.Coverage() < r.minCoverage
}
