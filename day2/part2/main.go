package main

import (
	"advent-of-code/shared"
	"fmt"
	"log"
	"strings"
)

func main() {
	scanner := shared.NewScanner("./input.txt")
	depth, forward, aim := 0, 0, 0

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		v := shared.StringToInt(parts[1])

		switch parts[0] {
		case "forward":
			forward += v
			depth += aim * v
		case "down":
			aim += v
		case "up":
			aim -= v
		default:
			log.Fatalf("failed to handle command: %v", parts[0])
		}
	}

	fmt.Println(depth * forward)
}
