package main

import (
	"fmt"
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

	unmatched := make(map[byte]int)
	for s, q := range target {
		unmatched[s] = q
	}

	for i := 0; i < lenWord; i++ {
		s := seq[i]
		if _, ok := target[s]; ok {
			unmatched[s]--
			if unmatched[s] == 0 {
				delete(unmatched, s)
			}
		}
	}

	count := 0
	if len(unmatched) == 0 {
		count++
	}
	for i := lenWord; i < lenSeq; i++ {
		currSymbol, prevSymbol := seq[i], seq[i-lenWord]

		if currSymbol != prevSymbol {
			if _, ok := target[currSymbol]; ok {
				unmatched[currSymbol]--
				if unmatched[currSymbol] == 0 {
					delete(unmatched, currSymbol)
				}
			}

			if _, ok := target[prevSymbol]; ok {
				unmatched[prevSymbol]++
				if unmatched[prevSymbol] == 0 {
					delete(unmatched, prevSymbol)
				}
			}
		}

		if len(unmatched) == 0 {
			count++
		}
	}
	fmt.Println(count)
}
