package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type products map[string]int
type customers map[string]products

func (db *customers) addPurchase(purchase []string) {
	c, p := purchase[0], purchase[1]
	q, _ := strconv.Atoi(purchase[2])

	if _, ok := (*db)[c]; !ok {
		(*db)[c] = make(products)
	}
	(*db)[c][p] += q
}

func readPurchases(db *customers) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		purchase := strings.Fields(sc.Text())
		db.addPurchase(purchase)
	}
}

func (db *customers) printTotal() error {
	customersSorted := make([]string, 0, len(*db))
	for c := range *db {
		customersSorted = append(customersSorted, c)
	}
	sort.Strings(customersSorted)

	wr := bufio.NewWriter(os.Stdout)
	for _, c := range customersSorted {
		_, err := wr.WriteString(fmt.Sprintf("%s:\n", c))
		if err != nil {
			return err
		}

		productsSorted := make([]string, 0, len((*db)[c]))
		for p := range (*db)[c] {
			productsSorted = append(productsSorted, p)
		}
		sort.Strings(productsSorted)

		for _, p := range productsSorted {
			_, err := wr.WriteString(fmt.Sprintf("%s %d\n", p, (*db)[c][p]))
			if err != nil {
				return err
			}
		}
	}
	return wr.Flush()
}

func main() {
	db := make(customers)
	readPurchases(&db)
	err := db.printTotal()
	if err != nil {
		return
	}
}
