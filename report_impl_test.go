package coverreportparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeCoveredFile struct{}

func (fakeCoveredFile) Add(Line) {}

func (fakeCoveredFile) Coverage() float64 { return 100 }

func (fakeCoveredFile) Info() (count, covered float64) { return 1, 1 }

type fakeUncoveredFile struct{}

func (fakeUncoveredFile) Add(Line) {}

func (fakeUncoveredFile) Coverage() float64 { return 0 }

func (fakeUncoveredFile) Info() (count, covered float64) { return 3, 0 }

func TestReport_Add(t *testing.T) {
	r := report{}

	line1 := fakeLine{}
	line2 := fakeLine{}

	_, ok := r.files["file1"]
	assert.False(t, ok)
	_, ok = r.files["file2"]
	assert.False(t, ok)

	r.Add("file1", &line1)
	_, ok = r.files["file1"]
	assert.True(t, ok)
	_, ok = r.files["file2"]
	assert.False(t, ok)

	r.Add("file2", &line2)
	_, ok = r.files["file1"]
	assert.True(t, ok)
	_, ok = r.files["file2"]
	assert.True(t, ok)
}

func TestReport_Coverage(t *testing.T) {
	r := report{
		files: map[string]File{
			"file1": fakeCoveredFile{},
		},
	}
	assert.Equal(t, 100.0, r.Coverage())

	r = report{
		files: map[string]File{
			"file2": fakeUncoveredFile{},
		},
	}
	assert.Equal(t, 0.0, r.Coverage())

	r = report{
		files: map[string]File{
			"file1": fakeCoveredFile{},
			"file2": fakeUncoveredFile{},
		},
	}
	assert.Equal(t, 25.0, r.Coverage())
}

func TestReport_String(t *testing.T) {
	r := report{
		files: map[string]File{
			"file1": fakeCoveredFile{},
		},
	}
	expected := `100.00%: 	file1
`
	assert.Equal(t, expected, r.String())

	r = report{
		files: map[string]File{
			"file2": fakeUncoveredFile{},
		},
	}
	expected = `0.00%: 	file2
`
	assert.Equal(t, expected, r.String())

	r = report{
		files: map[string]File{
			"file1": fakeCoveredFile{},
			"file2": fakeUncoveredFile{},
		},
	}
	expected = `100.00%: 	file1
0.00%: 	file2
`
	assert.Equal(t, expected, r.String())
}
