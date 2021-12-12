package main

import (
	"advent-of-code/shared"
	"fmt"
)

type octopus struct {
	flashedDuringStep bool
	energyLevel int
	x int
	y int
}

const gridSize = 10

var grid = make(map[int]map[int]*octopus)
var totalFlashes = 0

func main()  {
	s := shared.NewScanner("./input.txt")

	y := 0
	for s.Scan() {
		grid[y] = map[int]*octopus{}
		for x, v := range s.Text() {
			grid[y][x] = &octopus{
				energyLevel: shared.StringToInt(string(v)),
				x: x,
				y: y,
			}
		}
		y++
	}

	for s := 0; s < 100; s++ {
		for y := 0; y < gridSize; y++ {
			for x := 0; x < gridSize; x++ {
				grid[y][x].flashedDuringStep = false
			}
		}

		for y := 0; y < gridSize; y++ {
			for x := 0; x < gridSize; x++ {
				grid[y][x].cycle()
			}
		}
	}
	fmt.Println(totalFlashes)
}

func (o *octopus) cycle() {
	if o.flashedDuringStep {
		return
	}

	o.energyLevel++
	if o.energyLevel <= 9 {
		return
	}

	o.flashedDuringStep = true
	o.energyLevel = 0
	totalFlashes++

	if o.y != 0 {
		grid[o.y-1][o.x].cycle()
	}
	if o.y != gridSize-1 {
		grid[o.y+1][o.x].cycle()
	}
	if o.x != 0 {
		grid[o.y][o.x-1].cycle()
	}
	if o.x != gridSize-1 {
		grid[o.y][o.x+1].cycle()
	}
	if o.y != 0 && o.x != 0 {
		grid[o.y-1][o.x-1].cycle()
	}
	if o.y != 0 && o.x != gridSize-1 {
		grid[o.y-1][o.x+1].cycle()
	}
	if o.y != gridSize-1 && o.x != 0 {
		grid[o.y+1][o.x-1].cycle()
	}
	if o.y != gridSize-1 && o.x != gridSize-1 {
		grid[o.y+1][o.x+1].cycle()
	}
}

