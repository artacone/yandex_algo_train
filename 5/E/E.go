package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())

	trees := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		trees[i], _ = strconv.Atoi(sc.Text())
	}

	counter := make(map[int]int)
	best := [2]int{0, n}
	for l, r := 0, 0; r < n; r++ {
		color := trees[r]
		if _, ok := counter[color]; !ok {
			counter[color] = 1
		} else {
			counter[color]++
		}

		if len(counter) == k {
			for counter[trees[l]] > 1 {
				counter[trees[l]]--
				l++
			}
			if r-l < best[1]-best[0] {
				best[0], best[1] = l, r
			}
		}
	}
	fmt.Println(best[0]+1, best[1]+1)
}
