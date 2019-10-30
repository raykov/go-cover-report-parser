package main

import (
	"strings"
	"testing"
)

var covered = []string{
	`mode: atomic
file1:1.1,2.10 1 1
file2:1.1,2.10 1 1
file3:1.1,2.10 1 1
file3:3.1,5.10 1 1
`,
	`mode: set
file1:1.1,2.10 1 10
file2:1.1,2.10 1 100
file3:1.1,2.10 1 1000
file3:3.1,5.10 1 2
`,
}

var uncovered = []string{
	`mode: atomic
file1:1.1,2.10 1 0
file2:1.1,2.10 1 1
file3:1.1,2.10 1 1
file3:3.1,5.10 1 1
`,
	`mode: set
file1:1.1,2.10 1 10
file2:1.1,2.10 1 100
file3:1.1,2.10 1 1000
file3:3.1,5.10 1 0
`,
}

func TestCheckCoverage(t *testing.T) {
	for _, data := range covered {
		cov, err := parseCoverage(strings.NewReader(data))
		if err != nil {
			t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n  for\n%v\n\n", err, nil, data)
		}
		coverage := cov.Coverage()
		if coverage != 100.0 {
			t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n  for\n%v\n\n", coverage, 100, data)
		}
	}

	for _, data := range uncovered {
		cov, err := parseCoverage(strings.NewReader(data))
		if err != nil {
			t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n  for\n%v\n\n", err, nil, data)
		}
		coverage := cov.Coverage()
		if coverage != 75.0 {
			t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n  for\n%v\n\n", coverage, 75, data)
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