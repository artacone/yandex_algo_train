package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Insert(n int) error {
	if t == nil {
		return fmt.Errorf("tree is nil")
	}
	if t.Value == n {
		return fmt.Errorf("already in tree: %d", n)
	} else if n < t.Value {
		if t.Left == nil {
			t.Left = &TreeNode{Value: n}
			return nil
		}
		return t.Left.Insert(n)
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Value: n}
			return nil
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
		root = &TreeNode{Value: n}
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

func (t *TreeNode) printForks(wr *bufio.Writer) {
	if t != nil {
		t.Left.printForks(wr)
		if t.Left != nil && t.Right != nil {
			wr.WriteString(strconv.Itoa(t.Value))
			wr.WriteByte('\n')
		}
		t.Right.printForks(wr)
	}
}

func main() {
	root := buildTree()
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	root.printForks(wr)
}
