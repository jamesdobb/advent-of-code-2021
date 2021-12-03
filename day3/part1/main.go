package main

import (
	"advent-of-code/shared"
	"fmt"
	"strconv"
)

func main() {
	scanner := shared.NewScanner("./input.txt")
	var t, f = make(map[int]int), make(map[int]int)

	for scanner.Scan() {
		v := scanner.Text()
		for i, c := range v {
			if fmt.Sprintf("%c", c) == "1" {
				t[i]++
			} else {
				f[i]++
			}
		}
	}

	gr, er := "", ""
	for i := 0; i < len(t); i++ {
		if t[i] == f[i] {
			panic("I don't know what to do here!")
		} else if t[i] > f[i] {
			gr += "1"
			er += "0"
		} else {
			gr += "0"
			er += "1"
		}
	}

	grV, _ := strconv.ParseInt(gr, 2, 64)
	erV, _ := strconv.ParseInt(er, 2, 64)


	fmt.Println(grV)
	fmt.Println(erV)

	fmt.Println(grV*erV)
}
