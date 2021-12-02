package shared

import (
	"log"
	"strconv"
)

func StringToInt(i string) int {
	v, err := strconv.Atoi(i)
	if err != nil {
		log.Fatalf("failed to parse value: %v", err)
	}
	return v
}