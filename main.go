package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("Usage: \n\n\t%s path_to_coverage.out \n\n", os.Args[0])
	}

	coverReportFilePath := os.Args[1]

	f, err := os.Open(coverReportFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	r := bufio.NewReader(f)

	// Skip first line
	//	mode: atomic
	r.ReadString('\n')

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			break
		}

		coverage := strings.Split(strings.TrimSpace(line), " ")

		if coverage[2] == "0" {
			log.Fatal("Some lines are not covered")
		}
	}
	log.Println("100% coverage")
}
