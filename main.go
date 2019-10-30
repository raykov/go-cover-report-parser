package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const usageMessage = "" +
	`Usage of 'go-jenkins-cover-report-parser':
Given a coverage profile produced by 'go test':
	go-jenkins-cover-report-parser -coverprofile=c.out

Set minimal expected coverage:
	go-jenkins-cover-report-parser -coverprofile=c.out -minimum=100
`

func usage() {
	fmt.Fprintln(os.Stderr, usageMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	minimumExpectedCoverage = flag.Float64("minimum", 100.00, "minimal expected coverage")
	coverprofile = flag.String("coverprofile", "", "a coverage profile produced by 'go test'")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	// Usage information when no arguments.
	if flag.NFlag() == 0 && flag.NArg() == 0 {
		flag.Usage()
	}

	if *minimumExpectedCoverage > 100.0 {
		*minimumExpectedCoverage = 100.0
	}

	f, err := os.Open(*coverprofile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	cov, err := parseCoverage(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	if cov.Coverage() < *minimumExpectedCoverage {
		log.Printf("Coverage (%.2f%%) is below the expected minimum coverage (%.2f%%).", cov.Coverage(), *minimumExpectedCoverage)
		os.Exit(2)
	}

	log.Printf("%.2f%% coverage\n", cov.Coverage())
}
