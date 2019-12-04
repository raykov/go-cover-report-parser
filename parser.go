package coverreportparser

import (
	"bufio"
	"io"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

// Parse parses provided reader and returns a coverage report
func Parse(reader interface{}) (Report, error) {
	var bufReader *bufio.Reader

	switch reader.(type) {
	case io.Reader:
		bufReader = bufio.NewReader(reader.(io.Reader))
	default:
		return nil, errors.New("reader should satisfy io.Reader interface")
	}

	// Skip first line
	//	mode: atomic
	bufReader.ReadString('\n')

	rep := report{}

	for {
		rawLine, err := bufReader.ReadString('\n')

		if err != nil {
			break
		}

		parsedLines, err := parseLine(rawLine)
		if err != nil {
			return nil, err
		}
		for _, parsedLine := range parsedLines {
			l := line{
				index:     parsedLine.index,
				statement: parsedLine.statement,
				covered:   parsedLine.covered,
			}

			rep.Add(parsedLine.fileName, &l)
		}
	}

	return &rep, nil
}

type parsedLine struct {
	fileName  string
	index     float64
	statement bool
	covered   bool
}

func parseLine(coverageLine string) ([]parsedLine, error) {
	reg := regexp.MustCompile("(.+):(\\d+)\\.(\\d+),(\\d+)\\.(\\d+)\\s(\\d+)\\s([01])")
	matches := reg.FindAllStringSubmatch(coverageLine, -1)
	if len(matches) == 0 || len(matches[0]) < 8 {
		return nil, errors.New("provided line doesn't match expected format")
	}

	_ = matches[0][0] // fullMatch
	fileName := matches[0][1]
	startLine, _ := strconv.ParseFloat(matches[0][2], 64)
	_, _ = strconv.ParseFloat(matches[0][3], 64) // startColumn
	endLine, _ := strconv.ParseFloat(matches[0][4], 64)
	_, _ = strconv.ParseFloat(matches[0][5], 64) // endColumn
	statementsCount, _ := strconv.ParseFloat(matches[0][6], 64)
	isCovered, _ := strconv.ParseFloat(matches[0][7], 64)

	parsedLines := make([]parsedLine, 0, int(endLine-startLine+1))

	for i := startLine; i < endLine+1; i++ {
		parsedLines = append(
			parsedLines,
			parsedLine{
				fileName:  fileName,
				index:     i,
				statement: statementsCount > 0,
				covered:   isCovered > 0,
			},
		)
	}

	return parsedLines, nil
}
