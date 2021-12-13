package main

import (
	"advent-of-code/shared"
	"fmt"
	"strings"
)

type fold struct {
	axis string
	pos  int
}

func main() {
	s := shared.NewScanner("./input.txt")

	codeMap := make(map[int]map[int]bool)

	for x := 0; x <= 2000; x++ {
		codeMap[x] = make(map[int]bool)
		for y := 0; y <= 2000; y++ {
			codeMap[x][y] = false
		}
	}

	maxX := 0
	maxY := 0
	for s.Scan() {
		l := s.Text()
		if len(l) == 0 {
			break
		}
		cords := strings.Split(l, ",")
		x := shared.StringToInt(cords[0])
		y := shared.StringToInt(cords[1])
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		codeMap[x][y] = true
	}

	var folds []fold
	for s.Scan() {
		f := strings.TrimLeft(s.Text(), "fold along ")
		p := strings.Split(f, "=")
		folds = append(folds, fold{p[0], shared.StringToInt(p[1])})
	}

	for _, f := range folds {
		startX := 0
		startY := 0
		if f.axis == "x" {
			startX = f.pos
		}
		if f.axis == "y" {
			startY = f.pos
		}

		for x := startX; x <= maxX; x++ {
			for y := startY; y <= maxY; y++ {
				if codeMap[x][y] == true {
					codeMap[x][y] = false
					if f.axis == "x" {
						newX := startX - (x - startX)
						codeMap[newX][y] = true
					} else {
						newY := startY - (y - startY)
						codeMap[x][newY] = true
					}
				}
			}
		}
	}

	dots := 0
	maxX2 := 0
	maxY2 := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if codeMap[x][y] == true {
				dots++
				if x > maxX2 {
					maxX2 = x
				}
				if y > maxY2 {
					maxY2 = y
				}
			}
		}
	}

	for y := 0; y <= maxY2; y++ {
		fmt.Print("\n")
		for x := 0; x <= maxX2; x++ {
			if codeMap[x][y] == true {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
	}
}
