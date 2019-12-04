package coverreportparser

// Check checks report and creates a result
func Check(report Report, minCoverage float64, verbose bool) Result {
	if minCoverage > 100.0 || minCoverage < 0.0 {
		minCoverage = 100.0
	}

	return &result{
		report:      report,
		minCoverage: minCoverage,
		verbose:     verbose,
	}
}
