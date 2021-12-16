package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type accounts map[string]int

func (db *accounts) add(account string, quantity int) {
	(*db)[account] += quantity
}

func (db *accounts) sub(account string, quantity int) {
	(*db)[account] -= quantity
}

func (db *accounts) transfer(from, to string, quantity int) {
	db.sub(from, quantity)
	db.add(to, quantity)
}

func (db *accounts) showBalance(account string) {
	if quantity, ok := (*db)[account]; ok {
		fmt.Println(quantity)
	} else {
		fmt.Println("ERROR")
	}
}

func (db *accounts) chargeIncome(p int) {
	k := 1. + float64(p)/100.
	for account, quantity := range *db {
		if quantity > 0 {
			(*db)[account] = int(float64(quantity) * k)
		}
	}

}

func main() {
	db := make(accounts)
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		query := strings.Fields(sc.Text())
		switch query[0] {
		case "DEPOSIT":
			name := query[1]
			quantity, _ := strconv.Atoi(query[2])
			db.add(name, quantity)
		case "WITHDRAW":
			name := query[1]
			quantity, _ := strconv.Atoi(query[2])
			db.sub(name, quantity)
		case "TRANSFER":
			from, to := query[1], query[2]
			quantity, _ := strconv.Atoi(query[3])
			db.transfer(from, to, quantity)
		case "BALANCE":
			name := query[1]
			db.showBalance(name)
		case "INCOME":
			p, _ := strconv.Atoi(query[1])
			db.chargeIncome(p)
		}
	}
}
