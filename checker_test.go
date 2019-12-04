package coverreportparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var coveredLine = line{
	index:     0,
	statement: true,
	covered:   true,
}
var uncoveredLine = line{
	index:     0,
	statement: true,
	covered:   false,
}

func TestCovered100Percents(t *testing.T) {
	report := report{
		files: map[string]File{
			"file1": &file{
				lines: []Line{
					&coveredLine,
				},
			},
			"file2": &file{
				lines: []Line{
					&coveredLine,
				},
			},
		},
	}
	expected := `100.00% coverage
`

	result := Check(&report, 9, false)

	assert.Equal(t, expected, result.String())
}

func TestCovered50Percents(t *testing.T) {
	report := report{
		files: map[string]File{
			"file1": &file{
				lines: []Line{
					&coveredLine,
				},
			},
			"file2": &file{
				lines: []Line{
					&uncoveredLine,
				},
			},
		},
	}
	expected := `50.00% coverage
`

	result := Check(&report, 0, false)

	assert.Equal(t, expected, result.String())
}

func TestCovered0Percents(t *testing.T) {
	report := report{
		files: map[string]File{
			"file1": &file{
				lines: []Line{
					&uncoveredLine,
				},
			},
			"file2": &file{
				lines: []Line{
					&uncoveredLine,
				},
			},
		},
	}
	expected := `0.00% coverage
`

	result := Check(&report, 0, false)

	assert.Equal(t, expected, result.String())
}

func TestUnCovered100Percents(t *testing.T) {
	report := report{
		files: map[string]File{
			"file1": &file{
				lines: []Line{
					&uncoveredLine,
				},
			},
			"file2": &file{
				lines: []Line{
					&uncoveredLine,
				},
			},
		},
	}
	expected := `Coverage (0.00%) is below the expected minimum coverage (100.00%).
`

	result := Check(&report, 100, false)

	assert.Equal(t, expected, result.String())
}

func TestUnCovered50Percents(t *testing.T) {
	report := report{
		files: map[string]File{
			"file1": &file{
				lines: []Line{
					&uncoveredLine,
				},
			},
			"file2": &file{
				lines: []Line{
					&uncoveredLine,
				},
			},
		},
	}
	expected := `Coverage (0.00%) is below the expected minimum coverage (50.00%).
`

	result := Check(&report, 50, false)

	assert.Equal(t, expected, result.String())
}

func TestVerbose(t *testing.T) {
	report := report{
		files: map[string]File{
			"file1": &file{
				lines: []Line{
					&uncoveredLine,
					&coveredLine,
				},
			},
			"file2": &file{
				lines: []Line{
					&uncoveredLine,
				},
			},
		},
	}
	expected := `50.00%: 	file1
0.00%: 	file2
Coverage (33.33%) is below the expected minimum coverage (100.00%).
`

	result := Check(&report, 100, true)

	assert.Equal(t, expected, result.String())
}

func TestMinExpectedCoverageBiggerThenOneHundred(t *testing.T) {
	report := report{
		files: map[string]File{
			"file1": &file{
				lines: []Line{
					&coveredLine,
				},
			},
			"file2": &file{
				lines: []Line{
					&coveredLine,
				},
			},
		},
	}
	expected := `100.00% coverage
`

	result := Check(&report, 1_000, false)

	assert.Equal(t, expected, result.String())
}

func TestMinExpectedCoverageBelowZero(t *testing.T) {
	report := report{
		files: map[string]File{
			"file1": &file{
				lines: []Line{
					&coveredLine,
				},
			},
			"file2": &file{
				lines: []Line{
					&coveredLine,
				},
			},
		},
	}
	expected := `100.00% coverage
`

	result := Check(&report, -1, false)

	assert.Equal(t, expected, result.String())
}
