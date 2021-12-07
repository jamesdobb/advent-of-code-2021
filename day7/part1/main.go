package main

import (
	"advent-of-code/shared"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main()  {
	s := shared.NewScanner("./input.txt")
	_ = s.Scan()
	l := s.Text()
	values := strings.Split(l, ",")

	var positions []int
	max := 0
	for _, sv := range values {
		v, _ := strconv.Atoi(sv)
		if v > max {
			max = v
		}
		positions = append(positions, v)
	}

	var fuelRequired = make(map[int]int)
	min := -1
	optimalIndex := 0
	for i := 0; i <= max; i++ {
		fuelRequired[i] = 0
		for _, p := range positions {
			fuelRequired[i] += int(math.Abs(float64(p - i)))
		}
		if fuelRequired[i] < min || min == -1 {
			min = fuelRequired[i]
			optimalIndex = i
		}
	}

	fmt.Println(fuelRequired[optimalIndex])

}
