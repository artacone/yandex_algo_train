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
	nBlocks, _ := strconv.Atoi(sc.Text())

	heights := make(map[int]int, nBlocks)
	for i := 0; i < nBlocks; i++ {
		sc.Scan()
		w, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		h, _ := strconv.Atoi(sc.Text())
		if h > heights[w] {
			heights[w] = h
		}
	}
	pyramidH := 0
	for _, h := range heights {
		pyramidH += h
	}
	fmt.Println(pyramidH)
}
