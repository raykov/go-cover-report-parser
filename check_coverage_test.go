package main

func ExampleCovered100Percents() {
	cov := coverage{
		"file1":{lines: 1, covered: 1},
	}

	checkCoverage(cov)
	// Output: 100.00% coverage
}

func ExampleCovered50Percents() {
	cov := coverage{
		"file1":{lines: 2, covered: 1},
	}

	*minimumExpectedCoverage = 50

	checkCoverage(cov)
	// Output: 50.00% coverage
}

func ExampleCovered0Percents() {
	cov := coverage{
		"file1":{lines: 1, covered: 0},
	}

	*minimumExpectedCoverage = 0

	checkCoverage(cov)
	// Output: 0.00% coverage
}

func ExampleUnCovered100Percents() {
	cov := coverage{
		"file1":{lines: 1, covered: 0},
	}

	*minimumExpectedCoverage = 100

	checkCoverage(cov)
	// Coverage (0.00%) is below the expected minimum coverage (100.00%).
}

func ExampleUnCovered50Percents() {
	cov := coverage{
		"file1":{lines: 2, covered: 0},
	}

	*minimumExpectedCoverage = 50

	checkCoverage(cov)
	// Coverage (0.00%) is below the expected minimum coverage (50.00%).
}

func ExampleUnCovered0Percents() {
	cov := coverage{
		"file1":{lines: 1, covered: 0},
	}

	*minimumExpectedCoverage = 0

	checkCoverage(cov)
	// Output: 0.00% coverage
}

func ExampleVerbose() {
	cov := coverage{
		"file1":{lines: 1, covered: 0},
		"file2":{lines: 100, covered: 50},
		"file3":{lines: 10, covered: 10},
		"file4":{lines: 100, covered: 75},
	}

	*minimumExpectedCoverage = 100
	*verbose = true

	checkCoverage(cov)
	// Output: 0.00%: 	file1
	//50.00%: 	file2
	//100.00%: 	file3
	//75.00%: 	file4
	//Coverage (63.98%) is below the expected minimum coverage (100.00%).
}
