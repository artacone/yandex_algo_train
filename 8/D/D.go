package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func buildTree() (root *TreeNode) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	if n == 0 {
		return
	} else {
		root = &TreeNode{Value: n, Height: 1}
	}
	for sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
		if n == 0 {
			break
		}
		root.Insert(n)
	}
	return
}

func (t *TreeNode) printInOrder() {
	if t != nil {
		t.Left.printInOrder()
		fmt.Println(t.Value)
		t.Right.printInOrder()
	}
}

func main() {
	root := buildTree()
	root.printInOrder()
}
