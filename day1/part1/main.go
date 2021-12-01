package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main()  {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}

	lastReading := 0
	increases := 0
	lines := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("failed to convert line to int: %v", err)
		}

		if v > lastReading && lastReading != 0 {
			increases++
		}

		lines++
		lastReading = v
	}

	fmt.Println(fmt.Sprintf("processed %v lines", lines))
	fmt.Println(fmt.Sprintf("result: %v", increases))
}
