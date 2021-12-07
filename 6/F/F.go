package main

import "fmt"

func getTimeToCopy(n, x, y int) int {
	if x > y {
		x, y = y, x
	}
	l, r := 0, n*x
	for l < r {
		m := int(uint(l+r) >> 1)
		if m/x+m/y >= n-1 {
			r = m
		} else {
			l = m + 1
		}
	}
	return l + x
}

func main() {
	var n, x, y int
	_, err := fmt.Scan(&n, &x, &y)
	if err != nil {
		return
	}
	fmt.Println(getTimeToCopy(n, x, y))
}
