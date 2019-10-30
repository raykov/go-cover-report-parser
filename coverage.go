package main

import (
	"fmt"
	"sort"
)

type coverage map[string]fileCov

func (c coverage) Coverage() float64 {
	var lines, covered float64

	for _, f := range c {
		lines += f.lines
		covered += f.covered
	}

	return covered / lines * 100
}

type keys []string

func (k keys) Len() int {
	return len(k)
}

func (k keys) Less(i, j int) bool {
	return k[i] < k[j]
}

func (k keys) Swap(i,j int) {
	k[i], k[j] = k[j], k[i]
}

func (c coverage) Print() {
	keys := make(keys, 0, len(c))
	for key, _ := range c {
		keys = append(keys, key)
	}

	sort.Sort(keys)

	for _, key := range keys {
		cov := c[key]
		fmt.Printf("%.2f%%: \t%s\n", cov.Coverage(), key)
	}
}
