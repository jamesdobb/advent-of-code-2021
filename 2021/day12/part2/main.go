package main

import (
	"advent-of-code/shared"
	"fmt"
	"strings"
)

type cave struct {
	id          string
	connections []string
}

var caves = map[string]*cave{}

func main() {
	s := shared.NewScanner("/Users/james/apps/advent-of-code/2021/day12/part1/input.txt")
	for s.Scan() {
		parts := strings.Split(s.Text(), "-")
		if _, ok := caves[parts[0]]; !ok {
			caves[parts[0]] = &cave{parts[0], []string{}}
		}
		if _, ok := caves[parts[1]]; !ok {
			caves[parts[1]] = &cave{parts[1], []string{}}
		}

		caves[parts[0]].connections = append(caves[parts[0]].connections, parts[1])
		caves[parts[1]].connections = append(caves[parts[1]].connections, parts[0])
	}

	fmt.Println(len(caves["start"].calcRoutes("")))
}

func (c *cave) calcRoutes(i string) []string {
	var r []string

	if c.id == "end" {
		return []string{i + ",end"}
	}

	if shared.IsLower(c.id) && c.id != "start" && c.id != "end" {
		if strings.Count(i, ","+c.id) == 1 {
			if string(i[0]) == "_"{
				return r
			}
			i = "_" + i
		} else if strings.Count(i, ","+c.id) == 2 {
			return r
		}
	}

	for _, v := range c.connections {
		lc := caves[v]
		if lc.id == "start" {
			continue
		}

		r = append(r, lc.calcRoutes(strings.TrimLeft(i+","+c.id, ","))...)
	}

	return r
}
