package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printNums(nums []int) {
	wNums := make([]string, len(nums))
	for i, n := range nums {
		wNums[i] = strconv.Itoa(n)
	}
	fmt.Println(strings.Join(wNums, " "))
}

func main() {
	counter := make(map[string]int)
	ans := make([]int, 0)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		word := s.Text()
		ans = append(ans, counter[word])
		counter[word]++
	}
	printNums(ans)
}
