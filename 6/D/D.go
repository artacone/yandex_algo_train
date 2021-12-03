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

func findDefSize(n, a, b, w, h int) (d int) {
	l, r := 0, w+1
	for l < r {
		m := int(uint(l+r+1) >> 1)
		if (w/(a+m+m))*(h/(b+m+m)) >= n {
			l = m
		} else {
			r = m - 1
		}
	}
	d = l
	l, r = 0, w+1
	for l < r {
		m := int(uint(l+r+1) >> 1)
		if (w/(b+m+m))*(h/(a+m+m)) >= n {
			l = m
		} else {
			r = m - 1
		}
	}
	d = max(d, l)
	return
}

func main() {
	var n, a, b, w, h int
	_, err := fmt.Scan(&n, &a, &b, &w, &h)
	if err != nil {
		log.Panic()
	}
	fmt.Println(findDefSize(n, a, b, w, h))
}
