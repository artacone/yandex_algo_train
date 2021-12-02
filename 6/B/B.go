package main

import (
	"bufio"
	"os"
	"strconv"
)

func abs(a int) int {
	if a < 0 {
		return 0 - a
	} else {
		return a
	}
}

func approximateSearch(n int, nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		middle := int(uint(left+right) >> 1)
		if nums[middle] == n {
			return n
		} else if nums[middle] < n {
			left = middle + 1
		} else {
			right = middle
		}
	}
	if left > 0 && abs(nums[left-1]-n) <= abs(nums[left]-n) {
		return nums[left-1]
	} else {
		return nums[left]
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		nums[i], _ = strconv.Atoi(sc.Text())
	}

	wr := bufio.NewWriter(os.Stdout)
	for i := 0; i < k; i++ {
		sc.Scan()
		num, _ := strconv.Atoi(sc.Text())
		closest := approximateSearch(num, nums)
		wr.WriteString(strconv.Itoa(closest))
		wr.WriteByte('\n')
	}
	wr.Flush()
}
