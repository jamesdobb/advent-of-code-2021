package main

import (
	"advent-of-code/shared"
	"fmt"
	"strings"
)

type display struct {
	inputs        map[string][]int // input string with slice of possible values (to be reduced to 1)
	displayValues []string
}

// primary index = number we want to match
// values = number to match against: number of segment matches expected
var match = map[int]map[int]int{
	0: {1: 2, 4: 3, 7: 3},
	2: {1: 1, 4: 2, 7: 2},
	3: {1: 2, 4: 3, 7: 3},
	5: {1: 1, 4: 3, 7: 2},
	6: {1: 1, 4: 3, 7: 2},
	9: {1: 2, 4: 4, 7: 3},
}

var lenPossibilities = map[int][]int{2: {1}, 3: {7}, 4: {4}, 5: {2, 3, 5}, 6: {0, 6, 9}, 7: {8}}

func main() {
	s := shared.NewScanner("./input.txt")

	t := 0
	for s.Scan() {
		p := strings.Split(s.Text(), "|")
		d := newDisplay(
			strings.Split(strings.Trim(p[0], " "), " "),
			strings.Split(strings.Trim(p[1], " "), " "),
		)
		t += d.out()
	}

	fmt.Println(t)
}

func newDisplay(inputs []string, displayValues []string) *display {
	var i = make(map[string][]int)
	for _, iv := range inputs {
		i[shared.SortString(iv)] = lenPossibilities[len(iv)]
	}

	var d []string
	for _, dv := range displayValues {
		d = append(d, shared.SortString(dv))
	}

	disp := &display{i, d}
	for s, vals := range disp.inputs {
		if len(vals) > 1{
			disp.inputs[s] = disp.reducePossibilities(s, vals)
		}
	}

	return disp
}

func (d *display) reducePossibilities(s string, vals []int) []int {
	for _, v := range vals {
		for _, ik := range []int{1, 4, 7} {
			if len(vals) == 1{
				break
			}

			exp, act := match[v][ik], 0

			for _, c := range d.known(ik) {
				if strings.Contains(s, string(c)) {
					act++
				}
			}

			if exp != act {
				vals = shared.Remove(vals, v)
			}
		}
	}

	return vals
}

func (d *display) out() int {
	out := ""

	for _, dv := range d.displayValues {
		out += fmt.Sprintf("%v", d.inputs[dv][0])
	}

	return shared.StringToInt(out)
}

func (d *display) known(i int) string {
	for s, j := range d.inputs {
		if len(j) == 1 && j[0] == i {
			return s
		}
	}

	return ""
}
