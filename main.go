package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		parts := strings.Split(os.Args[0], "/")
		log.Fatalf("Usage: \n\n\t%s path_to_coverage.out \n\n", parts[len(parts) - 1])
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
}

func checkCoverage(reader interface{}) (err error) {
	var r *bufio.Reader

	switch reader.(type) {
	case io.Reader:
		r = bufio.NewReader(reader.(io.Reader))
	default:
		return errors.New("reader should satisfy io.Reader interface")
	}

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
			return errors.New("some lines are not covered")
		}
	}

	log.Println("100% coverage")
	return nil
}
