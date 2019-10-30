package main

import (
	"strings"
	"testing"
)

var covered = []string{
	`mode: atomic
file1:1.1,2.10 1 1
file2:1.1,2.10 2 1
file3:1.1,2.10 2 1
file3:3.1,5.10 3 1
`,
	`mode: set
file1:1.1,20.10 20 10
file2:1.1,2.10 1 100
file3:1.1,3.10 3 1000
file3:1.5,12.10 7 2
`,
}
var coveredExpectedResults = []float64{100, 100}

var uncovered = []string{
	`mode: atomic
file1:1.1,2.10 1 0
file2:1.1,2.10 33 1
file3:1.1,2.10 33 1
file3:3.1,5.10 33 1
`,
	`mode: set
file1:1.1,2.10 1 10
file2:1.1,2.10 1 100
file3:1.1,2.10 1 1000
file3:3.1,5.10 97 0
`,
}
var uncoveredExpectedResults = []float64{99, 3}

func TestCheckCoverage(t *testing.T) {
	for i, data := range covered {
		cov, err := parseCoverage(strings.NewReader(data))
		if err != nil {
			t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n  for\n%v\n\n", err, nil, data)
		}
		coverage := cov.Coverage()
		if coverage != coveredExpectedResults[i] {
			t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n  for\n%v\n\n", coverage, coveredExpectedResults[i], data)
		}
	}

	for i, data := range uncovered {
		cov, err := parseCoverage(strings.NewReader(data))
		if err != nil {
			t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n  for\n%v\n\n", err, nil, data)
		}
		coverage := cov.Coverage()
		if coverage != uncoveredExpectedResults[i] {
			t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n  for\n%v\n\n", coverage, uncoveredExpectedResults[i], data)
		}
	}
}

func TestCheckCoverageNotAReader(t *testing.T) {
	_, err := parseCoverage("")
	if err != notAReaderError {
		t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n\n", err, notAReaderError)
	}

	_, err = parseCoverage(strings.NewReader(covered[0]))
	if err != nil {
		t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n\n", err, nil)
	}
}
