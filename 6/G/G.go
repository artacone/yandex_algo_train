package main

import "fmt"

func getBorderWidth(width, length, t int64) int64 {
	if width > length {
		width, length = length, width
	}
	perimeter := width + length
	left, right := int64(0), width/2
	for left < right {
		middle := int64(uint(left+right+1) >> 1)
		dmiddle := middle << 1
		tiles := (perimeter - dmiddle) * dmiddle
		if tiles > t || dmiddle > width {
			right = middle - 1
		} else {
			left = middle
		}
	}
	return left
}

func main() {
	var n, m, t int64
	_, err := fmt.Scan(&n, &m, &t)
	if err != nil {
		return
	}
	fmt.Println(getBorderWidth(n, m, t))
}
