package main

import (
	"fmt"
	"sort"
)

func main() {
	var brick [3]int
	var hole [2]int
	fmt.Scan(&brick[0], &brick[1], &brick[2], &hole[0], &hole[1])

	sort.Ints(brick[:])
	sort.Ints(hole[:])

	if brick[0] <= hole[0] && brick[1] <= hole[1] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
