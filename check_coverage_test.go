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
		err := checkCoverage(strings.NewReader(data))
		if err != nil {
			t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n  for\n%v\n\n", err, nil, data)
		}
	}

	for _, data := range uncovered {
		err := checkCoverage(strings.NewReader(data))
		if err != uncoveredError {
			t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n  for\n%v\n\n", err, uncoveredError, data)
		}
	}
}


func TestCheckCoverageNotAReader(t *testing.T) {
	err := checkCoverage("")
	if err != notAReaderError {
		t.Errorf("returned unexpected result: \ngot  %v \nwant %v\n\n", err, notAReaderError)
	}
}
