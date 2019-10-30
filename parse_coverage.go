package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var uncoveredError = errors.New("some lines are not covered")
var notAReaderError = errors.New("reader should satisfy io.Reader interface")

func parseCoverage(reader interface{}) (cov coverage, err error) {
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
		lines, _ := strconv.ParseFloat(parsedLine[1], 64)

		if _, ok := cov[fileName]; !ok {
			cov[fileName] = fileCov{}
		}
		fCov := cov[fileName]

		if parsedLine[2] == "0" {
			fCov.Add(lines, 0)
			//return uncoveredError
		} else {
			fCov.Add(lines, lines)
		}

		cov[fileName] = fCov
	}

	return cov, nil
}
