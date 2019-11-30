package coverreportparser

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var notAReaderError = errors.New("reader should satisfy io.Reader interface")

func parseReport(reader interface{}) (cov coverage, err error) {
	var r *bufio.Reader

	switch reader.(type) {
	case io.Reader:
		r = bufio.NewReader(reader.(io.Reader))
	default:
		return cov, notAReaderError
	}

	// Skip first line
	//	mode: atomic
	r.ReadString('\n')

	if cov == nil {
		cov = coverage{}
	}

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			break
		}

		parsedLine := strings.Split(strings.TrimSpace(line), " ")
		fileName := strings.Split(parsedLine[0], ":")[0]
		statementsCount, err := strconv.ParseFloat(parsedLine[1], 64)
		if err != nil {
			return nil, err
		}

		if _, ok := cov[fileName]; !ok {
			cov[fileName] = fileCoverage{}
		}
		fCov := cov[fileName]

		switch {
		case parsedLine[2] == "0":
			fCov.Add(statementsCount, 0)
		case parsedLine[1] == "0":
			/*
				The case when we have no statements in the code,
				but covered function with the tests.

					type SomeInterface interface {
						DoSomething()
					}

					type fakeStruct struct{}

					func (fakeStruct) DoSomething() {}

				fakeStruct satisfies SomeInterface and provides fake functionality
				for test purposes, for instance. But it does nothing.
				If you will cover this code with tests, in cover report you will get
				a line like

					file.go 0 1

				In this case we need to skip this line.
				I just want to have this explicit.
			*/
		default:
			fCov.Add(statementsCount, statementsCount)
		}

		cov[fileName] = fCov
	}

	return cov, nil
}
