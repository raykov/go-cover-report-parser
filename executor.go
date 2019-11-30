package coverreportparser

import (
	"errors"
	"os"
)

type Options struct {
	CoverProfile            string
	Verbose                 bool
	MinimumExpectedCoverage float64
}

var UncoveredError = errors.New("some lines are not covered")

func Execute(opts Options) error {
	reader, err := os.Open(opts.CoverProfile)
	if err != nil {
		return err
	}
	defer reader.Close()

	cov, err := parse(reader)
	if err != nil {
		return err
	}

	err = check(cov, opts)
	if err != nil {
		return err
	}

	return nil
}
