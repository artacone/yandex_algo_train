package main

import "fmt"

func getVasyaPlace(scores []int, n int) int {
	bestScore := scores[0]
	afterWinner := true
	vasyaScore := 0
	for i := 1; i < n-1; i++ {
		score := scores[i]
		if score > bestScore {
			bestScore = score
			vasyaScore = 0
			afterWinner = true
		} else if afterWinner {
			if score%10 == 5 && scores[i+1] < score && vasyaScore < score {
				vasyaScore = score
			}
		}
	}
	if vasyaScore == 0 {
		return 0
	}
	vasyaPlace := 1
	for i := 0; i < n; i++ {
		if scores[i] > vasyaScore {
			vasyaPlace += 1
		}
	}
	return vasyaPlace
}

func main() {
	var n int
	fmt.Scan(&n)
	scores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&(scores[i]))
	}
	fmt.Println(getVasyaPlace(scores, n))
}
