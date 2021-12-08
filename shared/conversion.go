package shared

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func StringToInt(i string) int {
	v, err := strconv.Atoi(i)
	if err != nil {
		log.Fatalf("failed to parse value: %v", err)
	}
	return v
}

func Remove(s []int, v int) []int {
	var r []int
	for _, i := range s {
		if i != v {
			r = append(r, i)
		}
	}

	return r
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}