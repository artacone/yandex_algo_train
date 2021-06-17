package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var a, b, n, m int
	fmt.Scan(&a, &b, &n, &m)

	aMin := n + (n-1)*a // n stops and at least (n - 1) intervals
	aMax := aMin + 2*a  // 2 more intervals
	bMin := m + (m-1)*b
	bMax := bMin + 2*b
	resMin := max(aMin, bMin)
	resMax := min(aMax, bMax)

	if resMin > resMax {
		fmt.Println(-1)
	} else {
		fmt.Println(resMin, resMax)
	}
}
