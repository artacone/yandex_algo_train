package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	dict := make(map[string]string, n+n)
	for i := 0; i < n; i++ {
		s.Scan()
		kv := strings.Fields(s.Text())
		k, v := kv[0], kv[1]
		dict[k] = v
		dict[v] = k
	}
	s.Scan()
	fmt.Println(dict[s.Text()])
}
