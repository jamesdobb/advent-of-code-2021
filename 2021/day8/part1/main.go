package main

import (
	"advent-of-code/shared"
	"fmt"
	"strings"
)

func main()  {
	s := shared.NewScanner("./input.txt")

	r := 0
	for s.Scan() {
		p := strings.Split(s.Text(), "|")

		l := strings.Trim(p[1], " ")
		for _, i := range strings.Split(l, " ") {
			if len(i) == 7 || len(i) == 4 || len(i) == 2 || len(i) == 3{
				r++
			}
		}
	}

	fmt.Println(r)
}
