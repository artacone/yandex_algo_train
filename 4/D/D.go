package main

import (
	"bufio"
	"os"
	"strconv"
)

func getDurability(sc *bufio.Scanner) map[int]int {
	sc.Scan()
	nKeys, _ := strconv.Atoi(sc.Text())
	durs := make(map[int]int, nKeys)

	for i := 1; i <= nKeys; i++ {
		sc.Scan()
		d, _ := strconv.Atoi(sc.Text())
		durs[i] = d
	}
	return durs
}

func useKeyboard(sc *bufio.Scanner, keyDur *map[int]int) {
	sc.Scan()
	nPressings, _ := strconv.Atoi(sc.Text())

	for i := 0; i < nPressings; i++ {
		sc.Scan()
		key, _ := strconv.Atoi(sc.Text())
		(*keyDur)[key]--
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	keyDurabilities := getDurability(sc)

	useKeyboard(sc, &keyDurabilities)

	results := make([]bool, len(keyDurabilities))
	for key, d := range keyDurabilities {
		results[key-1] = d < 0
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, isBroken := range results {
		if isBroken {
			wr.WriteString("YES\n")
		} else {
			wr.WriteString("NO\n")
		}
	}
	wr.Flush()
}
