package main

import "fmt"

func main() {
	count := 0
	var prev, curr, next int
	scan, _ := fmt.Scanf("%d %d %d", &prev, &curr, &next)
	if scan != 3 {
		fmt.Println(0)
		return
	}
	for {
		if prev < curr && curr > next {
			count++
		}
		prev = curr
		curr = next
		scan, _ := fmt.Scanf("%d", &next)
		if scan == 0 {
			break
		}
	}
	fmt.Println(count)
}
