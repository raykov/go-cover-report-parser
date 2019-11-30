package coverreportparser

import (
	"fmt"
	"sort"
)

type coverage map[string]fileCoverage

func (c coverage) Coverage() float64 {
	var lines, covered float64

	for _, f := range c {
		lines += f.lines
		covered += f.covered
	}

	return covered / lines * 100
}

type keys []string

func (c coverage) Print() {
	keys := make(keys, 0, len(c))
	for key, _ := range c {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, key := range keys {
		cov := c[key]
		fmt.Printf("%.2f%%: \t%s\n", cov.Coverage(), key)
	}
}
