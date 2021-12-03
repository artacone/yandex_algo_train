package main

import (
	"fmt"
	"log"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findBoardSide(w, h, n int) (m int) {
	l := max(w, h)
	r := l * n
	for l < r {
		m = int(uint(l+r) >> 1)
		if n <= (m/w)*(m/h) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	var w, h, n int
	_, err := fmt.Scan(&w, &h, &n)
	if err != nil {
		log.Panic()
	}
	fmt.Println(findBoardSide(w, h, n))
}
