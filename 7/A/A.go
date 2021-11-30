package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	START = -1
	END   = 1
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	nStudents, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	nTeachers, _ := strconv.Atoi(sc.Text())
	events := make([][2]int, 0, nTeachers+nTeachers)
	for i := 0; i < nTeachers; i++ {
		sc.Scan()
		l, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		r, _ := strconv.Atoi(sc.Text())

		events = append(events, [2]int{l, START})
		events = append(events, [2]int{r, END})

	}

	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})

	nWatchers := 0
	leftWatched := 0
	nGooglers := nStudents
	for _, event := range events {
		desk, eventType := event[0], event[1]
		switch eventType {
		case START:
			if nWatchers == 0 {
				leftWatched = desk
			}
			nWatchers++
		case END:
			nWatchers--
			if nWatchers == 0 {
				nGooglers -= desk - leftWatched + 1
			}
		}
	}
	fmt.Println(nGooglers)
}
