package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		lenWord, lenSeq int
		word, seq       string
	)
	_, err := fmt.Scanf("%d %d\n%s\n%s\n", &lenWord, &lenSeq, &word, &seq)
	if err != nil {
		return
	}

	target := make(map[byte]int, lenWord)
	for i := range word {
		target[word[i]]++
	}

	current := make(map[byte]int)
	for i := 0; i < lenWord; i++ {
		current[seq[i]]++
	}

	count := 0
	if reflect.DeepEqual(current, target) {
		count++
	}
	for i := lenWord; i < lenSeq; i++ {
		currSymbol, prevSymbol := seq[i], seq[i-lenWord]
		current[currSymbol]++
		current[prevSymbol]--
		if current[prevSymbol] < 1 {
			delete(current, prevSymbol)
		}
		if reflect.DeepEqual(current, target) {
			count++
		}
	}
	fmt.Println(count)
}
