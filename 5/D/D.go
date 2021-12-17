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
	d, _ := strconv.Atoi(sc.Text())

	monuments := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		monuments[i], _ = strconv.Atoi(sc.Text())
	}

	nWays := 0
	for l, r := 0, 1; r < n; {
		if monuments[r]-monuments[l] > d {
			nWays += n - r
			l++
		} else {
			r++
		}
	}
	fmt.Println(nWays)
}
