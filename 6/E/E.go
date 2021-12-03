package main

import "fmt"

func findNFives(a, b, c int) int {
	l, r := 0, a+b+c
	prevSum := 2*a + 3*b + 4*c
	prevN := a + b + c
	for l < r {
		m := int(uint(l+r) >> 1)
		if 2*(prevSum+5*m) >= 7*(prevN+m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		return
	}
	fmt.Println(findNFives(a, b, c))
}
