package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type void struct{}

type set map[int]void

var inSet = void{}

func readColors(sc *bufio.Scanner) (colorsA, colorsB set) {
	sc.Scan()
	nA, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	nB, _ := strconv.Atoi(sc.Text())

	colorsA = make(set)
	for i := 0; i < nA; i++ {
		sc.Scan()
		c, _ := strconv.Atoi(sc.Text())
		colorsA[c] = inSet
	}

	colorsB = make(set)
	for i := 0; i < nB; i++ {
		sc.Scan()
		c, _ := strconv.Atoi(sc.Text())
		colorsB[c] = inSet
	}
	return
}

func (s1 *set) getSetIntersection(s2 *set) []int {
	intersection := make([]int, 0, len(*s1))
	for n := range *s1 {
		if _, ok := (*s2)[n]; ok {
			intersection = append(intersection, n)
		}
	}
	sort.Ints(intersection)
	return intersection
}

func (s1 *set) getSetDiff(s2 *set) []int {
	diff := make([]int, 0, len(*s1))
	for n := range *s1 {
		if _, ok := (*s2)[n]; !ok {
			diff = append(diff, n)
		}
	}
	sort.Ints(diff)
	return diff
}

func printNums(nums []int) {
	wNums := make([]string, len(nums))
	for i, n := range nums {
		wNums[i] = strconv.Itoa(n)
	}
	fmt.Println(strings.Join(wNums, " "))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	colorsA, colorsB := readColors(sc)

	intersection := colorsA.getSetIntersection(&colorsB)
	fmt.Println(len(intersection))
	printNums(intersection)

	aDiffB := colorsA.getSetDiff(&colorsB)
	fmt.Println(len(aDiffB))
	printNums(aDiffB)

	bDiffA := colorsB.getSetDiff(&colorsA)
	fmt.Println(len(bDiffA))
	printNums(bDiffA)
}
