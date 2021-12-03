package main

import "fmt"

func findNFives(a, b, c int) int {
	l, r := 0, a+b+c
	twicePrevSum := 4*a + 6*b + 8*c
	sevenPrevN := 7 * (a + b + c)
	for l < r {
		m := int(uint(l+r) >> 1)
		if twicePrevSum+3*m >= sevenPrevN { // (prevSum + 5*m)/(prevN+m) >= 3.5
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
