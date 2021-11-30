package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Value  int
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

func (t *TreeNode) Insert(n int) (h int) {
	if t == nil {
		log.Panic()
	}
	if t.Value == n {
		return
	} else if n < t.Value {
		if t.Left == nil {
			h = t.Height + 1
			t.Left = &TreeNode{Value: n, Height: h}
			return
		}
		return t.Left.Insert(n)
	} else {
		if t.Right == nil {
			h = t.Height + 1
			t.Right = &TreeNode{Value: n, Height: h}
			return
		}
		return t.Right.Insert(n)
	}
}

func buildTree() (root *TreeNode, heights []int) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	if n == 0 {
		return
	} else {
		heights = append(heights, 1)
		root = &TreeNode{Value: n, Height: 1}
	}
	for sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
		if n == 0 {
			break
		}
		heights = append(heights, root.Insert(n))
	}
	return
}

func printNums(nums []int) {
	wNums := make([]string, len(nums))
	for i, n := range nums {
		wNums[i] = strconv.Itoa(n)
	}
	fmt.Println(strings.Join(wNums, " "))
}

func main() {
	_, heights := buildTree()
	printNums(heights)
}
