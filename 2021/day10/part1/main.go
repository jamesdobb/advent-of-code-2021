package main

import (
	"advent-of-code/shared"
	"fmt"
)

var pairMap = map[string]string {
	"{" : "}",
	"[" : "]",
	"(" : ")",
	"<" : ">",
}

func main() {
	s := shared.NewScanner("./input.txt")

	points := 0

	for s.Scan() {
		l := s.Text()

		open := ""
		for i := 0; i < len(l); i++ {
			c := string(l[i])

			if c == "(" || c == "[" || c == "{" || c == "<" {
				open += string(l[i])
				continue
			}

			if len(open) == 0 {
				panic("no opening but found a close")
			}

			expClose := pairMap[string(open[len(open)-1])]
			if expClose != c {
				switch c {
				case ")": points += 3
				case "]": points += 57
				case "}": points += 1197
				case ">": points += 25137
				}
				break
			} else {
				open = open[:len(open)-1]
			}
		}
	}

	fmt.Println(points)
}