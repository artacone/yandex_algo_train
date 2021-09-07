package main

import (
	"fmt"
)

func main() {
	max := [3]int{-1000000, -1000000, -1000000}
	min := [2]int{1000000, 1000000}
	var n int
	for {
		scan, _ := fmt.Scanf("%d", &n)
		if scan == 0 {
			break
		}
		if n > max[2] {
			max[0], max[1], max[2] = max[1], max[2], n
		} else if n > max[1] {
			max[0], max[1] = max[1], n
		} else if n > max[0] {
			max[0] = n
		}
		if n < min[0] {
			min[0], min[1] = n, min[0]
		} else if n < min[1] {
			min[1] = n
		}
	}
	if min[0]*min[1]*max[2] > max[0]*max[1]*max[2] {
		fmt.Println(min[0], min[1], max[2])
	} else {
		fmt.Println(max[0], max[1], max[2])
	}
}
