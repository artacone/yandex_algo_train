package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counter := make(map[string]int)
	mostFrequent := struct {
		word string
		freq int
	}{}

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		word := sc.Text()
		counter[word]++
		freq := counter[word]
		if freq > mostFrequent.freq {
			mostFrequent.word = word
			mostFrequent.freq = freq
		} else if freq == mostFrequent.freq && strings.Compare(word, mostFrequent.word) == -1 {
			mostFrequent.word = word
		}
	}
	fmt.Println(mostFrequent.word)
}
