package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	setWords := make(map[string]struct{})

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		setWords[scanner.Text()] = struct{}{}
	}
	fmt.Println(len(setWords))
}
