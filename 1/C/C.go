package main

import (
	"fmt"
	"strings"
)

func formatNumber(number *string) {
	*number = strings.Replace(*number, "+7", "8", -1)
	*number = strings.Replace(*number, "(", "", -1)
	*number = strings.Replace(*number, ")", "", -1)
	*number = strings.Replace(*number, ")", "", -1)
	*number = strings.Replace(*number, "-", "", -1)
	if len(*number) == 7 {
		*number = "8495" + *number
	}

}

func main() {
	var newNumber, oldNumber string
	fmt.Scan(&newNumber)
	formatNumber(&newNumber)
	for i := 0; i < 3; i++ {
		fmt.Scan(&oldNumber)
		formatNumber(&oldNumber)
		if newNumber == oldNumber {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
