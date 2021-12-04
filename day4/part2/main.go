package main

import (
	"advent-of-code/shared"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const boardSize = 5

func main() {
	scanner := shared.NewScanner("./input.txt")
	_ = scanner.Scan()

	var drawNumbers []int
	drawStrings := strings.Split(scanner.Text(), ",")
	for _, v := range drawStrings {
		n, _ := strconv.Atoi(v)
		drawNumbers = append(drawNumbers, n)
	}

	scanner.Scan()
	var boards []board
	for {
		b, s := newBoard(scanner)
		if !s {
			break
		}

		boards = append(boards, b)
	}

	var winners = make(map[int]bool)
	for _, d := range drawNumbers {
		for i, b := range boards {
			b.match(d)
			if b.hasWon() {
				winners[i] = true

				if len(winners) == len(boards){
					fmt.Println(b.score(d))
					return
				}
			}
		}
	}
}

func (b board) score(d int) int {
	sum := 0

	for _, c := range b {
		if !c.matched {
			sum += c.val
		}
	}

	fmt.Println(sum)
	fmt.Println(d)

	return sum*d
}

func newBoard(scanner *bufio.Scanner) (board, bool) {
	var b board

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}

		for _, v := range strings.Split(scanner.Text(), " ") {
			if len(v) == 0 {
				continue
			}
			n, _ := strconv.Atoi(v)
			b = append(b, cell{n,false })
		}
	}

	return b, len(b) > 0
}

func (b board) match(d int)  {
	for i, _ := range b {
		if d == b[i].val {
			b[i].matched = true
		}
	}
}

func(b board) hasWon() bool {
	for c := 0; c < boardSize; c++ {
		sum := 0
		for x := 0; x < boardSize; x++ {
			if b[(x*boardSize) + c].matched {
				sum++
			}
			if sum == boardSize {
				return true
			}
		}
	}

	for r := 0; r < boardSize; r++ {
		sum := 0
		for i := 0; i < boardSize; i++ {
			if b[(r*boardSize) + i].matched {
				sum++
			}
			if sum == boardSize {
				return true
			}
		}
	}

	return false
}

type board []cell

type cell struct {
	val int
	matched bool
}
