package main

import (
	"fmt"
)

func main() {
	min, secondToMin := 1, 1
	max, secondToMax := -1, -1
	var n int
	for {
		scan, _ := fmt.Scanf("%d", &n)
		if scan == 0 {
			break
		}
		if n < 0 {
			if n < min {
				secondToMin, min = min, n
			} else if n < secondToMin {
				secondToMin = n
			}
		} else {
			if n > max {
				secondToMax, max = max, n
			} else if n > secondToMax {
				secondToMax = n
			}
		}
	}
	maxProduct, minProduct := max*secondToMax, min*secondToMin
	if maxProduct > minProduct {
		fmt.Println(secondToMax, max)
	} else {
		fmt.Println(min, secondToMin)
	}
}
