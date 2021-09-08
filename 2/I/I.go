package main

import "fmt"

func makeField(sizeX, sizeY, nMines int) [][]int {
	field := make([][]int, sizeY+2)
	for i := range field {
		field[i] = make([]int, sizeX+2)
	}

	var xMine, yMine int
	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < nMines; i++ {
		fmt.Scan(&yMine, &xMine)
		field[yMine][xMine] = -9
		for j := 0; j < 8; j++ {
			field[yMine+dy[j]][xMine+dx[j]] += 1
		}
	}
	return field
}

func printField(sizeX, sizeY int, field [][]int) {
	for y := 1; y <= sizeY; y++ {
		for x := 1; x <= sizeX; x++ {
			current := field[y][x]
			if current < 0 {
				fmt.Printf("* ")
			} else {
				fmt.Printf("%d ", current)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	var sizeX, sizeY, nMines int
	fmt.Scan(&sizeY, &sizeX, &nMines)

	field := makeField(sizeX, sizeY, nMines)
	printField(sizeX, sizeY, field)
}
