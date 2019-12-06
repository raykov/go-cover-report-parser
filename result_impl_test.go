package coverreportparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeSuccessReport struct{}

func (fakeSuccessReport) Add(string, Line)  {}
func (fakeSuccessReport) Coverage() float64 { return 100 }
func (fakeSuccessReport) String() string    { return "File1 is covered" }

type fakeFailureReport struct{}

func (fakeFailureReport) Add(string, Line)  {}
func (fakeFailureReport) Coverage() float64 { return 0 }
func (fakeFailureReport) String() string    { return "File1 is not covered" }

func TestResult_String(t *testing.T) {
	t.Run("when plain output", func(t *testing.T) {
		resultString(
			t,
			0,
			false,
			fakeSuccessReport{},
			"100.00% coverage\n",
		)
	})

	t.Run("when verbose output", func(t *testing.T) {
		resultString(
			t,
			0,
			true,
			fakeSuccessReport{},
			"File1 is covered100.00% coverage\n",
		)
	})

	t.Run("when success report", func(t *testing.T) {
		resultString(
			t,
			100,
			false,
			fakeSuccessReport{},
			"100.00% coverage\n",
		)
	})

	t.Run("when failure report", func(t *testing.T) {
		resultString(
			t,
			100,
			false,
			fakeFailureReport{},
			"Coverage (0.00%) is below the expected minimum coverage (100.00%).\n",
		)
	})
}

func resultString(t *testing.T, minCoverage float64, verbose bool, report Report, expected string) {
	r := result{
		minCoverage: minCoverage,
		verbose:     verbose,
		report:      report,
	}

	assert.Equal(t, expected, r.String())
}

func TestResult_IsFailure(t *testing.T) {
	t.Run("when success report", func(t *testing.T) {
		r := result{
			minCoverage: 100,
			report:      fakeSuccessReport{},
		}

		assert.False(t, r.IsFailure())
	})

	t.Run("when failure report", func(t *testing.T) {
		r := result{
			minCoverage: 100,
			report:      fakeFailureReport{},
		}

		assert.True(t, r.IsFailure())
	})
}
