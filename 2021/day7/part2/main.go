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

	var crabShipPositions = make(map[int]int)
	largestDistance := 0
	for _, shipPosStr := range values {
		shipPos, _ := strconv.Atoi(shipPosStr)
		if shipPos > largestDistance {
			largestDistance = shipPos
		}
		crabShipPositions[shipPos]++
	}

	var fuelRequired = make(map[int]int)
	minFuel := -1
	for i := 0; i <= largestDistance; i++ {
		for p, c := range crabShipPositions {
			diff := math.Abs(float64(p - i))
			fuelRequired[i] += int(diff * ((diff+1)/2)) * c
		}
		if fuelRequired[i] < minFuel || minFuel == -1 {
			minFuel = fuelRequired[i]
		}
	}

	fmt.Println(minFuel)
}
