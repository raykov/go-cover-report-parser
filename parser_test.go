package coverreportparser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("when information is unreadable", func(t *testing.T) {
		_, err := Parse("")
		assert.Equal(t, "reader should satisfy io.Reader interface", err.Error())
	})

	t.Run("when coverage file is invalid", func(t *testing.T) {
		data := "mode: atomic\nfile2:1.1,2.10 a 1\n"

		reader := strings.NewReader(data)
		_, err := Parse(reader)

		assert.Equal(t, "provided line doesn't match expected format", err.Error())
	})

	data := `mode: atomic
file1:1.1,2.10 1 1
file2:1.1,2.10 2 1
file3:1.1,2.10 2 1
file3:3.1,5.10 3 1
`
	reader := strings.NewReader(data)
	report, err := Parse(reader)
	assert.Nil(t, err)
	assert.Equal(t, 100.0, report.Coverage())

	data = `mode: set
file1:1.1,20.10 20 1
file2:1.1,2.10 1 1
file3:1.1,3.10 3 1
file3:1.5,12.10 7 1
`
	reader = strings.NewReader(data)
	report, _ = Parse(reader)
	assert.Equal(t, 100.0, report.Coverage())

	data = `mode: atomic
file1:1.1,1.10 1 0
file2:1.1,33.10 33 1
file3:1.1,33.10 33 1
file3:35.1,67.10 33 1
`
	reader = strings.NewReader(data)
	report, _ = Parse(reader)
	assert.Equal(t, 99.0, report.Coverage())

	data = `mode: set
file1:1.1,1.10 1 1
file2:1.1,1.10 1 1
file3:1.1,1.10 1 1
file3:3.1,99.10 97 0
`
	reader = strings.NewReader(data)
	report, _ = Parse(reader)
	assert.Equal(t, 3.0, report.Coverage())

	data = `mode: set
file1:1.1,1.10 0 1
file1:1.1,1.10 1 1
file2:1.1,1.10 1 0
file3:1.1,1.10 1 1
file3:3.1,3.10 1 1
`
	reader = strings.NewReader(data)
	report, _ = Parse(reader)
	assert.Equal(t, 75.0, report.Coverage())
}

func TestParseLine(t *testing.T) {
	parsedLines, err := parseLine("github.com/raykov/go-cover-report-parser/test_file_2.go:9.1,9.10 1 1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parsedLines))
	assert.Equal(t, "github.com/raykov/go-cover-report-parser/test_file_2.go", parsedLines[0].fileName)
	assert.Equal(t, 9.0, parsedLines[0].index)
	assert.True(t, parsedLines[0].covered)
	assert.True(t, parsedLines[0].statement)

	parsedLines, err = parseLine("test_file_2.go:9.1,9.10 0 1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parsedLines))
	assert.Equal(t, "test_file_2.go", parsedLines[0].fileName)
	assert.Equal(t, 9.0, parsedLines[0].index)
	assert.True(t, parsedLines[0].covered)
	assert.False(t, parsedLines[0].statement)

	parsedLines, err = parseLine("a:9.1,9.10 1 0")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(parsedLines))
	assert.Equal(t, "a", parsedLines[0].fileName)
	assert.Equal(t, 9.0, parsedLines[0].index)
	assert.False(t, parsedLines[0].covered)
	assert.True(t, parsedLines[0].statement)

	parsedLines, err = parseLine("github.com/raykov/go-cover-report-parser/test_file_2.go:9.1,10.10 1 1")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(parsedLines))
	assert.Equal(t, "github.com/raykov/go-cover-report-parser/test_file_2.go", parsedLines[0].fileName)
	assert.Equal(t, 9.0, parsedLines[0].index)
	assert.True(t, parsedLines[0].covered)
	assert.True(t, parsedLines[0].statement)
	assert.Equal(t, "github.com/raykov/go-cover-report-parser/test_file_2.go", parsedLines[1].fileName)
	assert.Equal(t, 10.0, parsedLines[1].index)
	assert.True(t, parsedLines[1].covered)
	assert.True(t, parsedLines[1].statement)

	_, err = parseLine(":9.1,9.10 1 1")
	assert.Equal(t, "provided line doesn't match expected format", err.Error())

	_, err = parseLine("github.com/raykov/go-cover-report-parser/test_file_2.go:a.1,9.10 1 1")
	assert.Equal(t, "provided line doesn't match expected format", err.Error())

	_, err = parseLine("github.com/raykov/go-cover-report-parser/test_file_2.go:9.b,9.10 1 1")
	assert.Equal(t, "provided line doesn't match expected format", err.Error())

	_, err = parseLine("github.com/raykov/go-cover-report-parser/test_file_2.go:9.1,c.10 1 1")
	assert.Equal(t, "provided line doesn't match expected format", err.Error())

	_, err = parseLine("github.com/raykov/go-cover-report-parser/test_file_2.go:9.1,9.d 1 1")
	assert.Equal(t, "provided line doesn't match expected format", err.Error())

	_, err = parseLine("github.com/raykov/go-cover-report-parser/test_file_2.go:9.1,9.10 e 1")
	assert.Equal(t, "provided line doesn't match expected format", err.Error())

	_, err = parseLine("github.com/raykov/go-cover-report-parser/test_file_2.go:9.1,9.10 1 f")
	assert.Equal(t, "provided line doesn't match expected format", err.Error())
}
