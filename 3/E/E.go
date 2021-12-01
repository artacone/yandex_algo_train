package main

import (
	"fmt"
	"sort"
)

type void struct{}

type set map[int]void

var inSet = void{}

func getRequiredKeys(n int) set {
	if n == 0 {
		return set{0: inSet}
	}
	required := make(set, 10)
	for n > 0 {
		digit := n % 10
		required[digit] = inSet
		n /= 10
	}
	return required
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

func main() {
	var x, y, z, n int
	_, err := fmt.Scanf("%d %d %d\n%d", &x, &y, &z, &n)
	if err != nil {
		return
	}
	keys := set{x: inSet, y: inSet, z: inSet}
	required := getRequiredKeys(n)

	fmt.Println(len(required.getSetDiff(&keys)))
}
