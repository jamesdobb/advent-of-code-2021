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
	depth := 0
	forward := 0

	for scanner.Scan() {
		command := scanner.Text()
		parts := strings.Split(command, " ")
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("failed to parse value: %v", err)
		}

		switch parts[0] {
		case "forward":
			forward = forward + v
		case "down":
			depth = depth + v
		case "up":
			depth = depth - v
		default:
			log.Fatalf("failed to handle command: %v", err)
		}
	}

	fmt.Println(depth * forward)
}
