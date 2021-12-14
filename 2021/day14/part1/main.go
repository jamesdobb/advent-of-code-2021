package main

import (
	"advent-of-code/shared"
	"fmt"
	"strings"
)

var rules = make(map[string]string)
var pairs = make(map[string]uint64)
var elements = make(map[string]uint64)

func main() {
	s := shared.NewScanner("./input.txt")
	s.Scan(); input := s.Text(); s.Scan()

	for s.Scan() { parts := strings.Split(s.Text(), " -> "); rules[parts[0]] = parts[1]; elements[parts[1]] = 0 }

	for _, c := range input {elements[string(c)]++}
	for i := 0; i < len(input)-1; i++ { pairs[string(input[i]) + string(input[i+1])]++ }

	for i := 0; i < 1000000; i++ {
		newPairs := make(map[string]uint64)
		for k, i := range pairs {
			if newChar, ok := rules[k]; ok {
				newPairs[string(k[0]) + newChar]+=i
				newPairs[newChar + string(k[1])]+=i
				elements[newChar] += i
			}
		}
		pairs = newPairs
	}

	min, max := uint64(0), uint64(0)
	for _, v := range elements {
		if v > max { max = v }; if v < min || min == 0 { min = v }
	}

	fmt.Println(max - min)
}
