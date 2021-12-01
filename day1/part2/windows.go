package main

import (
	"bufio"
	"log"
)

const windowSize = 3
type measurementWindow [windowSize]int

func NewWindows(s *bufio.Scanner) (*measurementWindow, *measurementWindow ) {
	w1 := &measurementWindow{}
	w2 := &measurementWindow{}

	for i := 0; i <= windowSize; i++ {
		v, success := nextMeasurement(s)
		if !success {
			log.Fatalf("failed to init window while scanning: %v", s.Err())
		}

		if i != windowSize {
			w1[i] = v
		}

		if i != 0 {
			w2[i-1] = v
		}
	}

	return w1, w2
}

func SlideWindows(s *bufio.Scanner, w1 *measurementWindow, w2 *measurementWindow) bool {
	v, success := nextMeasurement(s)
	if !success {
		return false
	}

	for i := 0; i < windowSize; i++ {
		w1[i] = w2[i]

		if i < windowSize-1 {
			w2[i] = w2[i+1]
		}
	}

	w2[windowSize-1] = v

	return true
}

func (mw *measurementWindow) total() int {
	t := 0

	for i := 0; i < windowSize; i++ {
		t = t + mw[i]
	}

	return t
}
