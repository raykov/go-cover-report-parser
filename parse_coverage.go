package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"
)

var uncoveredError = errors.New("some lines are not covered")
var notAReaderError = errors.New("reader should satisfy io.Reader interface")

type fileCov struct {
	lines float64
	covered float64
}

func (fc *fileCov) Add(lines, covered float64) {
	fc.lines += lines
	fc.covered += covered
}

func (fc *fileCov) Coverage() float64 {
	return fc.covered / fc.lines * 100
}

type coverage map[string]fileCov

func (c *coverage) Coverage() float64 {
	var lines, covered float64

	for _, f := range *c {
		lines += f.lines
		covered += f.covered
	}

	return covered / lines * 100
}

func (c *coverage) Print() {
	for n, f := range *c {
		log.Printf("%.2f%%: \t%s\n", f.Coverage(), n)
	}
}

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
