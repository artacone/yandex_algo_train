package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var tRoom, tCond, tResult int
	var mode string
	fmt.Scan(&tRoom, &tCond, &mode)

	switch mode {
	case "auto":
		tResult = tCond
	case "fan":
		tResult = tRoom
	case "freeze":
		tResult = min(tCond, tRoom)
	case "heat":
		tResult = max(tCond, tRoom)
	}
	fmt.Println(tResult)
}
