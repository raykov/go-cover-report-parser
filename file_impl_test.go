package coverreportparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeLine struct {
	statement bool
	covered   bool
}

func (f fakeLine) isCovered() bool {
	return f.covered
}

func (f fakeLine) isStatement() bool {
	return f.statement
}

func TestFile_Add(t *testing.T) {
	f := file{
		lines: []Line{},
	}

	assert.Equal(t, 0, len(f.lines))
	f.Add(fakeLine{})
	f.Add(fakeLine{})
	assert.Equal(t, 2, len(f.lines))
}

func TestFile_Coverage(t *testing.T) {
	uncoveredLine := fakeLine{statement: true}
	coveredLine := fakeLine{covered: true, statement: true}
	coveredFile := file{
		lines: []Line{
			coveredLine,
			coveredLine,
		},
	}
	uncoveredFile := file{
		lines: []Line{
			coveredLine,
			uncoveredLine,
		},
	}

	assert.Equal(t, 100.0, coveredFile.Coverage())
	assert.Equal(t, 50.0, uncoveredFile.Coverage())
}

func TestFile_Info(t *testing.T) {
	uncoveredLine := fakeLine{statement: true}
	coveredLine := fakeLine{covered: true, statement: true}

	uncoveredEmptyLine := fakeLine{statement: false}
	coveredEmptyLine := fakeLine{covered: true, statement: false}

	coveredFile := file{
		lines: []Line{
			coveredLine,
			coveredLine,
			coveredEmptyLine,
			uncoveredEmptyLine,
		},
	}
	uncoveredFile := file{
		lines: []Line{
			coveredLine,
			uncoveredLine,
			coveredEmptyLine,
			uncoveredEmptyLine,
		},
	}

	count, covered := coveredFile.Info()
	assert.Equal(t, 2.0, count)
	assert.Equal(t, 2.0, covered)

	count, covered = uncoveredFile.Info()
	assert.Equal(t, 2.0, count)
	assert.Equal(t, 1.0, covered)
}
