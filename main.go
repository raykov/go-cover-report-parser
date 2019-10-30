package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		parts := strings.Split(os.Args[0], "/")
		log.Fatalf("Usage: \n\n\t%s path_to_coverage.out \n\n", parts[len(parts)-1])
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	err = checkCoverage(f)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("100% coverage")
}
