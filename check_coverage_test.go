package coverreportparser

func ExampleCovered100Percents() {
	cov := coverage{
		"file1":{lines: 1, covered: 1},
	}
	opts := Options{
		CoverProfile:            "",
		Verbose:                 false,
		MinimumExpectedCoverage: 100,
	}

	check(cov, opts)
	// Output: 100.00% coverage
}

func ExampleCovered50Percents() {
	cov := coverage{
		"file1":{lines: 2, covered: 1},
	}
	opts := Options{
		CoverProfile:            "",
		Verbose:                 false,
		MinimumExpectedCoverage: 50,
	}


	check(cov, opts)
	// Output: 50.00% coverage
}

func ExampleCovered0Percents() {
	cov := coverage{
		"file1":{lines: 1, covered: 0},
	}
	opts := Options{
		CoverProfile:            "",
		Verbose:                 false,
		MinimumExpectedCoverage: 0,
	}

	check(cov, opts)
	// Output: 0.00% coverage
}

func ExampleUnCovered100Percents() {
	cov := coverage{
		"file1":{lines: 1, covered: 0},
	}

	opts := Options{
		CoverProfile:            "",
		Verbose:                 false,
		MinimumExpectedCoverage: 100,
	}

	check(cov, opts)
	// Output: Coverage (0.00%) is below the expected minimum coverage (100.00%).
}

func ExampleUnCovered50Percents() {
	cov := coverage{
		"file1":{lines: 2, covered: 0},
	}

	opts := Options{
		CoverProfile:            "",
		Verbose:                 false,
		MinimumExpectedCoverage: 50,
	}

	check(cov, opts)
	// Output: Coverage (0.00%) is below the expected minimum coverage (50.00%).
}

func ExampleUnCovered0Percents() {
	cov := coverage{
		"file1":{lines: 1, covered: 0},
	}

	opts := Options{
		CoverProfile:            "",
		Verbose:                 false,
		MinimumExpectedCoverage: 0,
	}

	check(cov, opts)
	// Output: 0.00% coverage
}

func ExampleVerbose() {
	cov := coverage{
		"file1":{lines: 1, covered: 0},
		"file2":{lines: 100, covered: 50},
		"file3":{lines: 10, covered: 10},
		"file4":{lines: 100, covered: 75},
	}

	opts := Options{
		CoverProfile:            "",
		Verbose:                 true,
		MinimumExpectedCoverage: 100,
	}

	check(cov, opts)
	// Output: 0.00%: 	file1
	//50.00%: 	file2
	//100.00%: 	file3
	//75.00%: 	file4
	//Coverage (63.98%) is below the expected minimum coverage (100.00%).
}

func ExampleMinExpectedCoverageBiggerThenOneHundred() {
	cov := coverage{
		"file1":{lines: 1, covered: 0},
	}

	opts := Options{
		CoverProfile:            "",
		Verbose:                 true,
		MinimumExpectedCoverage: 1_000,
	}

	check(cov, opts)
	// Output: 0.00%: 	file1
	//Coverage (0.00%) is below the expected minimum coverage (100.00%).
}

func ExampleMinExpectedCoverageBelowZero() {
	cov := coverage{
		"file1":{lines: 1, covered: 0},
	}

	opts := Options{
		CoverProfile:            "",
		Verbose:                 true,
		MinimumExpectedCoverage: -1,
	}

	check(cov, opts)
	// Output: 0.00%: 	file1
	//Coverage (0.00%) is below the expected minimum coverage (100.00%).
}
