package main

import (
	"advent-of-code/shared"
	"fmt"
	"strconv"
)

func main() {
	scanner := shared.NewScanner("./input.txt")
	inputs := shared.ScanStringSlice(scanner)

	ogr := calcRating(inputs, 0, "1", "0", "1")
	csr := calcRating(inputs, 0, "0", "1", "0")

	ogrV, _ := strconv.ParseInt(ogr, 2, 64)
	csrV, _ := strconv.ParseInt(csr, 2, 64)

	fmt.Println(ogrV * csrV)
}

func calcRating(s []string, i int, tV, fV, dV string) string {
	t, f := bitsAtPos(i, s)

	if t > f {
		s = filter(i, tV, s)
	} else if t < f {
		s = filter(i, fV, s)
	} else {
		s = filter(i, dV, s)
	}

	if len(s) == 1 {
		return s[0]
	}

	return calcRating(s, i+1, tV, fV, dV)
}

func filter(i int, m string, s []string) []string {
	var r []string

	for _, v := range s {
		if string(v[i]) == m {
			r = append(r, v)
		}
	}

	return r
}

func bitsAtPos(i int, s []string) (int, int) {
	r := map[string]int{"0": 0, "1": 0}

	for _, v := range s {
		r[string(v[i])]++
	}

	return r["1"], r["0"]
}
