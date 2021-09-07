package main

import (
	"fmt"
)

func main() {
	var prevNum, newNum int
	var sequenceType string
	print(sequenceType)
	fmt.Scan(&newNum)
	for {
		prevNum = newNum
		fmt.Scan(&newNum)
		if newNum == -2000000000 {
			break
		}
		if newNum == prevNum {
			switch sequenceType {
			case "":
				sequenceType = "CONSTANT"
			case "ASCENDING":
				sequenceType = "WEAKLY ASCENDING"
			case "DESCENDING":
				sequenceType = "WEAKLY DESCENDING"
			}
		} else if newNum > prevNum {
			switch sequenceType {
			case "":
				sequenceType = "ASCENDING"
			case "CONSTANT":
				sequenceType = "WEAKLY ASCENDING"
			case "WEAKLY DESCENDING", "DESCENDING":
				sequenceType = "RANDOM"
			}
		} else {
			switch sequenceType {
			case "":
				sequenceType = "DESCENDING"
			case "CONSTANT":
				sequenceType = "WEAKLY DESCENDING"
			case "WEAKLY ASCENDING", "ASCENDING":
				sequenceType = "RANDOM"
			}
		}
		if sequenceType == "RANDOM" {
			break
		}
	}
	fmt.Println(sequenceType)
}
