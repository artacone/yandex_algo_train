package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type void struct{}
type set map[int]void

var inSet = void{}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	s := make(set)
	for i := 0; i < n; i++ {
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		s[x] = inSet
		sc.Scan()
	}
	fmt.Println(len(s))
}
