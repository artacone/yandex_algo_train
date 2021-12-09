package main

import (
	"fmt"
)

type void struct{}
type genome string
type set map[genome]void

var inSet = void{}

func getPairs(g *genome) set {
	n := len(*g)
	pairs := make(set)
	for i := 0; i < n-1; i++ {
		pairs[(*g)[i:i+2]] = inSet
	}
	return pairs
}

func (g1 *genome) getProximity(g2 *genome) (proximity int) {
	pairs2 := getPairs(g2)
	n := len(*g1)

	for i := 0; i < n-1; i++ {
		if _, ok := pairs2[(*g1)[i:i+2]]; ok {
			proximity++
		}
	}
	return
}

func main() {
	var g1, g2 genome
	_, err := fmt.Scanf("%s\n%s", &g1, &g2)
	if err != nil {
		return
	}
	fmt.Println(g1.getProximity(&g2))
}
