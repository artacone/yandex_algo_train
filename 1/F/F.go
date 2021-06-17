package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


func newTable(w, l int) table {
	var t table
	t.w = w
	t.l = l
	t.area = w*l
	return t
}

type table struct {
	w, l, area int
}

func main()  {
	var w1, w2, l1, l2 int
	fmt.Scan(&w1, &l1, &w2, &l2)

	var tables [4]table
	tables[0] = newTable(w1 + w2, max(l1, l2))
	tables[1] = newTable(w1 + l2, max(l1, w2))
	tables[2] = newTable(max(w1, l2), l1 + w2)
	tables[3] = newTable(max(w1, w2), l1 + l2)

	res := tables[0]
	for i := 1; i < 4; i++ {
		if tables[i].area < res.area {
			res = tables[i]
		}
	}
	fmt.Println(res.w, res.l)
}