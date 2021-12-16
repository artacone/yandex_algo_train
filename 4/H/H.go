package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 4e6)
	sc.Buffer(buf, 4e6)
	sc.Scan()
	lens := strings.Fields(sc.Text())
	lenWord, _ := strconv.Atoi(lens[0])
	lenSeq, _ := strconv.Atoi(lens[1])
	sc.Scan()
	word := sc.Text()
	sc.Scan()
	seq := sc.Text()

	target := make(map[byte]int, 26+26)
	for i := range word {
		target[word[i]]++
	}

	unmatched := make(map[byte]int, len(target))
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
