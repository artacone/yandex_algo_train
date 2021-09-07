package main

import (
	"fmt"
)

func main() {
	min, secondToMin := 1000000, 1000000
	max, secondToMax := -1000000, -1000000
	var n, count int
	for {
		scan, _ := fmt.Scanf("%d", &n)
		if scan == 0 {
			break
		}
		if n < min {
			secondToMin, min = min, n
		} else if n < secondToMin {
			secondToMin = n
		}
		if n > max {
			secondToMax, max = max, n
		} else if n > secondToMax {
			secondToMax = n
		}
		count++
	}
	if count == 2 {
		fmt.Println(min, max)
		return
	}
	if max*secondToMax > min*secondToMin {
		fmt.Println(secondToMax, max)
	} else {
		fmt.Println(min, secondToMin)
	}
}
