package main

import "log"

type coverage map[string]fileCov

func (c *coverage) Coverage() float64 {
	var lines, covered float64

	for _, f := range *c {
		lines += f.lines
		covered += f.covered
	}

	return covered / lines * 100
}

func (c *coverage) Print() {
	for n, f := range *c {
		log.Printf("%.2f%%: \t%s\n", f.Coverage(), n)
	}
}
