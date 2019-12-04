package main

import (
	"flag"
	"fmt"
	"os"

	coverReport "github.com/raykov/go-cover-report-parser"
)

const usageMessage = "" +
	`Usage of 'go-jenkins-cover-report-parser':
Given a coverage profile produced by 'go test':
	go-jenkins-cover-report-parser -coverprofile=c.out

Set minimal expected coverage:
	go-jenkins-cover-report-parser -coverprofile=c.out -minimum=100`

func usage() {
	fmt.Fprintln(os.Stderr, usageMessage)
	fmt.Fprintln(os.Stderr, "Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	minimumExpectedCoverage = flag.Float64("minimum", 100.00, "minimal expected coverage")
	coverprofile            = flag.String("coverprofile", "", "a coverage profile produced by 'go test'")
	verbose                 = flag.Bool("verbose", false, "detailed coverage report")
	v                       = flag.Bool("v", false, "detailed coverage report")
)

func main() {
	setupFlags()

	reader, err := os.Open(*coverprofile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer reader.Close()

	report, err := coverReport.Parse(reader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	result := coverReport.Check(report, *minimumExpectedCoverage, *verbose || *v)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	if result.IsFailure() {
		os.Exit(1)
	}
}

func setupFlags() {
	flag.Usage = usage
	flag.Parse()

	// Usage information when no arguments.
	if flag.NFlag() == 0 && flag.NArg() == 0 {
		flag.Usage()
	}
}
