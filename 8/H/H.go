package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TreeNode struct {
	Value      int
	LeftDepth  int
	RightDepth int
	Left       *TreeNode
	Right      *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

func (t *TreeNode) Insert(n int) (err error) {
	if t == nil {
		return fmt.Errorf("tree is nil")
	}
	if t.Value == n {
		return fmt.Errorf("already in tree: %d", n)
	} else if n < t.Value {
		if t.Left == nil {
			t.Left = &TreeNode{Value: n}
			t.LeftDepth = 1
			return nil
		}
		err = t.Left.Insert(n)
		t.LeftDepth = 1 + max(t.Left.LeftDepth, t.Left.RightDepth)
		return
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Value: n}
			t.RightDepth = 1
			return nil
		}
		err = t.Right.Insert(n)
		t.RightDepth = 1 + max(t.Right.LeftDepth, t.Right.RightDepth)
		return err
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

func (t *TreeNode) isAVLBalanced() bool {
	stack, node := make([]*TreeNode, 1), t
	stack[0] = node
	for len(stack) > 0 {
		stack, node = stack[:len(stack)-1], stack[len(stack)-1]
		if abs(node.LeftDepth-node.RightDepth) > 1 {
			return false
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return true
}

func main() {
	root := buildTree()
	if root.isAVLBalanced() {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
