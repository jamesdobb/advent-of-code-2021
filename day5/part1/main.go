package main

import (
	"advent-of-code/shared"
	"fmt"
	"strconv"
	"strings"
)

const diagX = 1000
const diagY = 1000

func main()  {
	s := shared.NewScanner("./input.txt")

	var diagram [ diagX ][ diagY ] int

	for x := 0; x < diagX; x++ {
		for y := 0; y < diagY; y++ {
			diagram[x][y] = 0
		}
	}

	for s.Scan() {
		r := strings.Split(s.Text(), "->")
		x1, y1 := parseDimensions(r[0])
		x2, y2 := parseDimensions(r[1])

		// only horizontal / vertical for now
		if x1 != x2 && y1 != y2 {
			continue
		}

		x := x1
		y := y1
		for {
			diagram[x][y]++

			if x != x2 {
				if x1 < x2 {
					x++
				} else {
					x--
				}
			}

			if y != y2 {
				if y1 < y2 {
					y++
				} else {
					y--
				}
			}

			if x == x2 && y == y2 {
				diagram[x][y]++
				break
			}
		}
	}

	dangerZones := 0
	for x := 0; x < diagX; x++ {
		for y := 0; y < diagY; y++ {
			if diagram[y][x] >= 2 {
				dangerZones++
			}
		}
	}

	fmt.Println(dangerZones)
}

func parseDimensions(i string) (int, int) {
	parts := strings.Split(strings.Trim(i, " "), ",")

	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return x, y
}
