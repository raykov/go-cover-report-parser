package coverreportparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	t.Run("when file is missing", func(t *testing.T) {
		opts := Options {
			CoverProfile:            "/tmp/some-missing-dir/some-missing-file",
			Verbose:                 false,
			MinimumExpectedCoverage: 100,
		}

		err := Execute(opts)
		assert.Equal(t, "open /tmp/some-missing-dir/some-missing-file: no such file or directory", err.Error())
	})

	t.Run("when file is invalid", func(t *testing.T) {
		opts := Options {
			CoverProfile:            "./testdata/coverage_invalid.out",
			Verbose:                 false,
			MinimumExpectedCoverage: 100,
		}

		err := Execute(opts)
		assert.Equal(t, `strconv.ParseFloat: parsing "a": invalid syntax`, err.Error())
	})

	t.Run("when coverage rate is below expected", func(t *testing.T) {
		opts := Options {
			CoverProfile:            "./testdata/coverage_0.out",
			Verbose:                 false,
			MinimumExpectedCoverage: 100,
		}

		err := Execute(opts)
		assert.Equal(t, `some lines are not covered`, err.Error())
	})

	t.Run("when everything is covered", func(t *testing.T) {
		opts := Options {
			CoverProfile:            "./testdata/coverage_100.out",
			Verbose:                 false,
			MinimumExpectedCoverage: 100,
		}

		err := Execute(opts)
		assert.Nil(t, err)
	})
}
