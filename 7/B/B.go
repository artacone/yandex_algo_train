package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

const (
	BEGIN = -1
	POINT = 0
	END   = 1
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	nSegments, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	nPoints, _ := strconv.Atoi(sc.Text())
	events := make([][3]int, 0, nSegments+nSegments+nPoints)
	for i := 0; i < nSegments; i++ {
		sc.Scan()
		l, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		r, _ := strconv.Atoi(sc.Text())
		if l > r {
			l, r = r, l
		}
		events = append(events, [3]int{l, BEGIN, 0})
		events = append(events, [3]int{r, END, 0})
	}
	for i := 0; i < nPoints; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		events = append(events, [3]int{p, POINT, i})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i][0] == events[j][0] {
			return events[i][1] < events[j][1]
		}
		return events[i][0] < events[j][0]
	})

	currIn := 0
	result := make([]int, nPoints)
	for _, event := range events {
		_, pType, pI := event[0], event[1], event[2]
		switch pType {
		case POINT:
			result[pI] = currIn
		case BEGIN:
			currIn++
		case END:
			currIn--
		}
	}

	wr := bufio.NewWriter(os.Stdout)
	for _, i := range result {
		wr.WriteString(strconv.Itoa(i))
		wr.WriteByte(' ')
	}
	wr.WriteByte('\n')
	wr.Flush()
}
