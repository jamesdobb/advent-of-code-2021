package main

import (
	"advent-of-code/shared"
	"fmt"
	"strconv"
	"strings"
)

const newFishTimer = 8
const resetFishTimer = 6

func main()  {
	s := shared.NewScanner("./input.txt")
	_ = s.Scan()
	input := strings.Split(s.Text(), ",")

	var fish []int
	for _, i := range input {
		v, _ := strconv.Atoi(i)
		fish = append(fish, v)
	}

	for d := 1; d <= 80; d++ {
		for i, f := range fish {
			switch f {
			case 0:
				fish = append(fish, newFishTimer)
				fish[i] = resetFishTimer
			default:
				fish[i]--
			}
		}
	}

	fmt.Println(len(fish))
}
