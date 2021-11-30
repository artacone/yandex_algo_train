package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func mySearch(n int, nums []int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := int(uint(l+r) >> 1)
		if nums[m] == n {
			return true
		} else if nums[m] < n {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	s.Scan()
	m, _ := strconv.Atoi(s.Text())

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		nums[i], _ = strconv.Atoi(s.Text())
	}
	sort.Ints(nums)

	for i := 0; i < m; i++ {
		s.Scan()
		num, _ := strconv.Atoi(s.Text())
		if mySearch(num, nums) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
