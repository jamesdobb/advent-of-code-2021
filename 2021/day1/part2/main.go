package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main()  {
	s := createScanner()
	w1, w2 := NewWindows(s)

	increases := 0

	for {
		if w2.total() > w1.total() {
			increases++
		}

		if !SlideWindows(s, w1, w2) {
			break
		}
	}

	fmt.Println(increases)
}

func createScanner() *bufio.Scanner {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	return bufio.NewScanner(f)
}

func nextMeasurement(s *bufio.Scanner) (int, bool) {
	if !s.Scan() {
		return 0, false
	}

	v, err := strconv.Atoi(s.Text())
	if err != nil {
		log.Fatalf("failed to convert line to int: %v", err)
	}

	return v, true
}
