package main

import (
	"fmt"
	"math"
)

const epsilon = 0.000001

func getFreqRange() (float64, float64) {
	var minFreq, maxFreq float64 = 30, 4000
	var n int
	var previous, current, average float64
	fmt.Scan(&n, &previous)

	var msg string
	for i := 1; i < n; i++ {
		fmt.Scanf("%f %s", &current, &msg)
		if math.Abs(current-previous) < epsilon {
			continue
		}
		average = (current + previous) / 2
		if msg == "closer" {
			if current > previous {
				minFreq = math.Max(minFreq, average)
			} else {
				maxFreq = math.Min(maxFreq, average)
			}
		} else {
			if current < previous {
				minFreq = math.Max(minFreq, average)
			} else {
				maxFreq = math.Min(maxFreq, average)
			}
		}
		previous = current
	}
	return minFreq, maxFreq
}

func main() {
	fmt.Println(getFreqRange())
}
