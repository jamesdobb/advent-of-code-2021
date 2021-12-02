package main

import (
	"advent-of-code/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	scanner := shared.NewScanner("./input.txt")
	depth, forward, aim := 0, 0, 0

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("failed to parse value: %v", err)
		}

		switch parts[0] {
		case "forward":
			forward = forward + v
			depth = depth + (aim * v)
		case "down":
			aim = aim + v
		case "up":
			aim = aim - v
		default:
			log.Fatalf("failed to handle command: %v", parts[1])
		}
	}

	fmt.Println(depth * forward)
}
