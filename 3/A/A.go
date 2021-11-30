package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type void struct{}

var exist void

func main() {
	setNums := make(map[int64]void)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		n, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Panic(err)
		}
		setNums[n] = exist
	}
	fmt.Println(len(setNums))
}
