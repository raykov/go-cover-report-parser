package coverreportparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine_IsCovered(t *testing.T) {
	coveredLine := line{
		index:   0,
		covered: true,
	}
	uncoveredLine := line{
		index:   1,
		covered: false,
	}

	assert.Equal(t, true, coveredLine.isCovered())
	assert.Equal(t, false, uncoveredLine.isCovered())
}

func TestLine_IsStatement(t *testing.T) {
	statementLine := line{
		index:     0,
		statement: true,
	}
	emptyLine := line{
		index:     1,
		statement: false,
	}

	assert.Equal(t, true, statementLine.isStatement())
	assert.Equal(t, false, emptyLine.isStatement())
}
