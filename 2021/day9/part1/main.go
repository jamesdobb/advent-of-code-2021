package main

import (
	"advent-of-code/shared"
	"fmt"
)

func main() {
	s := shared.NewScanner("./input.txt")

	heights := make(map[int]map[int]int)

	row := 0
	for s.Scan() {
		heights[row] = make(map[int]int)
		line := s.Text()
		for i := 0; i < len(line); i++ {
			heights[row][i] = shared.StringToInt(string(line[i]))
		}
		row++
	}

	numRows := len(heights)
	numColumns := len(heights[0])

	riskLevels := 0
	for y := 0; y < numRows; y++ {
		for x := 0; x < numColumns; x++ {
			exp, act := 0, 0
			v := heights[y][x]

			if x != 0 {
				exp++
				if heights[y][x-1] > v {
					act++
				}
			}
			if x != numColumns-1 {
				exp++
				if heights[y][x+1] > v {
					act++
				}
			}
			if y != 0 {
				exp++
				if heights[y-1][x] > v {
					act++
				}
			}
			if y != numRows-1 {
				exp++
				if heights[y+1][x] > v {
					act++
				}
			}

			if act == exp {
				riskLevels += v + 1
			}
		}
	}

	fmt.Println(riskLevels)
}
