package main

import (
	"bufio"
	"os"
	"strconv"
)

type void struct{}
type set map[string]void

var inSet = void{}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	nStudents, _ := strconv.Atoi(sc.Text())

	langsEveryoneKnows := make(set)
	langs := make(set)
	for i := 0; i < nStudents; i++ {
		sc.Scan()
		nLangs, _ := strconv.Atoi(sc.Text())
		currentLangs := make(set, nLangs)
		for j := 0; j < nLangs; j++ {
			sc.Scan()
			langs[sc.Text()] = inSet
			currentLangs[sc.Text()] = inSet
		}
		if i == 0 {
			for lang := range currentLangs {
				langsEveryoneKnows[lang] = inSet
			}
		} else {
			for lang := range langsEveryoneKnows {
				if _, ok := currentLangs[lang]; !ok {
					delete(langsEveryoneKnows, lang)
				}
			}
		}
	}

	wr := bufio.NewWriter(os.Stdout)

	wr.WriteString(strconv.Itoa(len(langsEveryoneKnows)))
	wr.WriteByte('\n')

	for lang := range langsEveryoneKnows {
		wr.WriteString(lang)
		wr.WriteByte('\n')
	}

	wr.WriteString(strconv.Itoa(len(langs)))
	wr.WriteByte('\n')

	for lang := range langs {
		wr.WriteString(lang)
		wr.WriteByte('\n')
	}
	wr.Flush()
}
