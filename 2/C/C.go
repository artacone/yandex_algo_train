package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var sizeNums int
	fmt.Scan(&sizeNums)
	nums := make([]int, sizeNums)
	for i := 0; i < sizeNums; i++ {
		fmt.Scan(&(nums[i]))
	}
	var x, diff int
	fmt.Scan(&x)
	closest := nums[0]
	closestDiff := abs(x - closest)

	for i := 0; i < sizeNums; i++ {
		diff = abs(x - nums[i])
		if diff < closestDiff {
			closest, closestDiff = nums[i], diff
		}
	}
	fmt.Println(closest)
}
