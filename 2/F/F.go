package main

import "fmt"

func isSymmetric(nums []int) bool {
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}
	return true
}

func reverseSlice(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func makeSymmetric(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if isSymmetric(nums[i:]) {
			return reverseSlice(nums[:i])
		}
	}
	return []int{}
}

func main() {
	var numsSize int
	fmt.Scan(&numsSize)
	nums := make([]int, numsSize)
	for i := 0; i < numsSize; i++ {
		fmt.Scan(&(nums[i]))
	}
	result := makeSymmetric(nums)
	fmt.Println(len(result))
	if len(result) > 0 {
		for _, n := range result {
			fmt.Printf("%d ", n)
		}
	}
}
