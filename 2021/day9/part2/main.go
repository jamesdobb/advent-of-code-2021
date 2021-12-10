package main

import (
	"advent-of-code/shared"
	"fmt"
)

var heights = make(map[int]map[int]int)
var numRows int
var numColumns int

func main() {
	s := shared.NewScanner("./input.txt")

	row := 0
	for s.Scan() {
		heights[row] = make(map[int]int)
		line := s.Text()
		for i := 0; i < len(line); i++ {
			heights[row][i] = shared.StringToInt(string(line[i]))
		}
		row++
	}

	numRows = len(heights)
	numColumns = len(heights[0])

	i := 0
	var basins = make(map[int]*basin)
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
				b := &basin{x,y, false, 0, make(map[string]bool)}
				b.cSize = b.size(x, y)
				basins[i] = b
				i++
			}
		}
	}

	var use []*basin
	for {
		largest := &basin{startX: -1}
		largestIndex := 0

		for i, b := range basins {
			if b.isSelected {
				continue
			}
			if largest.startX == -1 || b.cSize > largest.cSize {
				largest = b
				largestIndex = i
			}
		}

		use = append(use, largest)
		if entry, ok := basins[largestIndex]; ok {
			entry.isSelected = true
			basins[largestIndex] = entry
		} else {
			panic ("fuck")
		}

		if len(use) == 3 {
			break
		}
	}

	r := 1
	for _, v := range use {
		fmt.Println(v.startX, v.startY, "|" , heights[v.startY][v.startX], "=", v.cSize)
		r = r * v.cSize
	}

	fmt.Println(r)
}

type basin struct {
	startX int
	startY int
	isSelected bool
	cSize int

	checkedLocations map[string]bool // "x.y"
}

func (b *basin) size(x int, y int) int {
	r := 0
	if b.checkedLocations[fmt.Sprintf("%.%", x, y)] {
		return r
	}
	b.checkedLocations[fmt.Sprintf("%.%", x, y)] = true

	v := heights[y][x]
	if v == 9 {
		return r
	}

	r++

	if y != 0 {
		r += b.size(x, y-1)
	}

	if y != numRows-1 {
		r += b.size(x, y+1)
	}

	if x != 0 {
		r += b.size(x-1, y)
	}

	if x != numColumns-1 {
		r += b.size(x+1, y)
	}

	return r
}
