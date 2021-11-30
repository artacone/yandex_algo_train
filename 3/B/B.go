package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type void struct{}

var exist void

func makeSet(r *bufio.Reader) map[int]void {
	set := make(map[int]void)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	words := strings.Fields(line)
	for _, w := range words {
		n, err := strconv.Atoi(w)
		if err != nil {
			log.Panic(err)
		}
		set[n] = exist
	}
	return set
}

func getSetIntersection(s1, s2 map[int]void) []int {
	intersection := make([]int, 0, len(s1))
	for n := range s1 {
		if _, ok := s2[n]; ok {
			intersection = append(intersection, n)
		}
	}
	sort.Ints(intersection)
	return intersection
}

func printNums(nums []int) {
	wNums := make([]string, len(nums))
	for i, n := range nums {
		wNums[i] = strconv.Itoa(n)
	}
	fmt.Println(strings.Join(wNums, " "))
}

func main() {
	r := bufio.NewReader(os.Stdin)
	set1, set2 := makeSet(r), makeSet(r)

	printNums(getSetIntersection(set1, set2))
}
