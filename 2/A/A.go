package main

import (
	"fmt"
)

func main() {
	var current, prev int
	fmt.Scan(&prev)
	for {
		scan, _ := fmt.Scan(&current)
		if scan == 0 {
			break
		}
		if current <= prev {
			fmt.Println("NO")
			return
		}
		prev = current
	}
	fmt.Println("YES")
}
