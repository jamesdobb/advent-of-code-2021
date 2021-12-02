package shared

import (
	"bufio"
	"log"
	"os"
)

func NewScanner(filename string) *bufio.Scanner {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	return bufio.NewScanner(f)
}

func ScanStringSlice(scanner *bufio.Scanner) []string {
	var r []string

	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	return r
}

func ScanIntSlice(scanner *bufio.Scanner) []int {
	 var r []int

	for scanner.Scan() {
		r = append(r, StringToInt(scanner.Text()))
	}

	return r
}
