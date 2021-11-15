package main

import "fmt"

func main() {
	setNums := make(map[int]bool)
	var n int
	for {
		scan, _ := fmt.Scanf("%d", &n)
		if scan == 0 {
			break
		}
		setNums[n] = true
	}
	println(len(setNums))
}
