package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	nCars, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	favourite, _ := strconv.Atoi(sc.Text())

	cars := make([]int, nCars)
	for i := 0; i < nCars; i++ {
		sc.Scan()
		cars[i], _ = strconv.Atoi(sc.Text())
	}

	count, currSum := 0, 0
	for left, right := 0, 0; right < nCars; right++ {
		currSum += cars[right]
		for currSum > favourite {
			currSum -= cars[left]
			left++
		}
		if currSum == favourite {
			count++
		}
	}
	fmt.Println(count)
}
