package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func readClothes(sc *bufio.Scanner) []int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	colors := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		colors[i], _ = strconv.Atoi(sc.Text())
	}
	return colors
}

func getStyle(shirts, pants []int) (int, int) {
	si, pi := 0, 0
	bestS, bestP := 0, 0
	leastDiff := math.MaxInt

	for si < len(shirts) && pi < len(pants) {
		currShirt, currPants := shirts[si], pants[pi]
		if currShirt == currPants {
			return currShirt, currPants
		}

		currDiff := currShirt - currPants
		if currDiff < 0 {
			currDiff = 0 - currDiff
		}
		if currDiff < leastDiff {
			bestS, bestP = si, pi
			leastDiff = currDiff
		}
		if currShirt > currPants {
			pi++
		} else {
			si++
		}
	}
	return shirts[bestS], pants[bestP]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	shirts := readClothes(sc)
	pants := readClothes(sc)
	fmt.Println(getStyle(shirts, pants))
}
