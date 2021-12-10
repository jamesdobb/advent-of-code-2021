package main

import (
	"advent-of-code/shared"
	"fmt"
	"sort"
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
	var pointsSlice []int

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
				open = ""
				break
			} else {
				open = open[:len(open)-1]
			}
		}

		if len(open) > 0 {
			incompletePoints := 0
			for j := len(open)-1; j >= 0; j-- {
				c := pairMap[string(open[j])]
				switch c {
					case ")": incompletePoints = (incompletePoints *5) + 1
					case "]": incompletePoints = (incompletePoints *5) + 2
					case "}": incompletePoints = (incompletePoints *5) + 3
					case ">": incompletePoints = (incompletePoints *5) + 4
				}
			}
			pointsSlice = append(pointsSlice, incompletePoints)
		}
	}

	sort.Ints(pointsSlice)

	fmt.Println(pointsSlice[len(pointsSlice)/2])

}