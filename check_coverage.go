package main

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func checkCoverage(reader interface{}) (err error) {
	var r *bufio.Reader

	switch reader.(type) {
	case io.Reader:
		r = bufio.NewReader(reader.(io.Reader))
	default:
		return errors.New("reader should satisfy io.Reader interface")
	}

	// Skip first line
	//	mode: atomic
	r.ReadString('\n')

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			break
		}

		coverage := strings.Split(strings.TrimSpace(line), " ")

		if coverage[2] == "0" {
			return errors.New("some lines are not covered")
		}
	}
	return nil
}
