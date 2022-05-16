package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type void struct{}

type turtle struct {
	ahead, behind int
}

type set map[turtle]void

var inSet = void{}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	nTurtles, _ := strconv.Atoi(sc.Text())

	turtles := make(set)
	for i := 0; i < nTurtles; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		if a+b == nTurtles-1 && a >= 0 && b >= 0 {
			t := turtle{ahead: a, behind: b}
			turtles[t] = inSet
		}
	}
	fmt.Println(len(turtles))
}
