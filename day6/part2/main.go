package main

import (
	"advent-of-code/shared"
	"fmt"
	"strconv"
	"strings"
)

const newFishTimer = 8
const resetFishTimer = 6
const days = 256

func main() {
	s := shared.NewScanner("./input.txt")
	_ = s.Scan()
	input := strings.Split(s.Text(), ",")

	var fish = make(map[int]int)
	for i := 0; i <= newFishTimer; i++ {
		fish[i] = 0
	}

	for _, i := range input {
		v, _ := strconv.Atoi(i)
		fish[v]++
	}

	for d := 0; d <= days; d++ {
		zerodFish := fish[0]

		for i := 1; i <= newFishTimer; i++ {
			fish[i-1] = fish[i]
		}

		fish[newFishTimer] = zerodFish
		fish[resetFishTimer] += zerodFish
	}

	t := 0
	for i := 0; i < newFishTimer; i++ {
		t += fish[i]
	}

	fmt.Println(t)
}
