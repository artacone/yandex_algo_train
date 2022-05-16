package main

import "fmt"

func longestSubstring(n, k int, in string) (int, int) {
	l, r := 0, 1
	maxLen, maxPos := 1, 1
	counter := make(map[byte]int)
	counter[in[l]] = 1
	for r < n {
		if counter[in[r]] == k {
			if r-l > maxLen {
				maxLen = r - l
				maxPos = l + 1
			}
			for counter[in[r]] == k {
				counter[in[l]]--
				l++
			}
		}
		counter[in[r]]++
		r++
	}
	if r-l > maxLen {
		maxLen = r - l
		maxPos = l + 1
	}
	return maxLen, maxPos
}

func main() {
	var (
		n     int
		k     int
		input string
	)
	_, err := fmt.Scanf("%d %d\n%s", &n, &k, &input)
	if err != nil {
		return
	}
	fmt.Println(longestSubstring(n, k, input))
}
