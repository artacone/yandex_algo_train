package main

import (
	"bufio"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scanPairs(sc *bufio.Scanner) [][2]int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	result := make([][2]int, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		result[i][0], _ = strconv.Atoi(sc.Text())
		sc.Scan()
		result[i][1], _ = strconv.Atoi(sc.Text())
	}

	return result
}

func buildPrefixUp(chain [][2]int) []int {
	n := len(chain)
	prefix := make([]int, n)
	for i := 1; i < n; i++ {
		up := max(chain[i][1]-chain[i-1][1], 0)
		prefix[i] = prefix[i-1] + up
	}
	return prefix
}

func buildPrefixUpBackwards(chain [][2]int) []int {
	n := len(chain)
	prefix := make([]int, n)
	for i := n - 1; i > 0; i-- {
		up := max(chain[i-1][1]-chain[i][1], 0)
		prefix[i-1] = prefix[i] + up
	}
	return prefix
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	chain := scanPairs(sc)
	routes := scanPairs(sc)

	prefixUp := buildPrefixUp(chain)
	prefixUpBackwards := buildPrefixUpBackwards(chain)

	wr := bufio.NewWriter(os.Stdout)
	for _, route := range routes {
		begin, end := route[0]-1, route[1]-1
		if begin <= end {
			up := prefixUp[end] - prefixUp[begin]
			wr.WriteString(strconv.Itoa(up))
		} else {
			up := prefixUpBackwards[end] - prefixUpBackwards[begin]
			wr.WriteString(strconv.Itoa(up))
		}
		wr.WriteByte('\n')
	}
	wr.Flush()
}
