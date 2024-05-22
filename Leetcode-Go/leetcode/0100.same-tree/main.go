package main

import (
	"fmt"

	"github.com/mo3et/leetcode-learn/leetcode-go/structures"
)

type TreeNode = structures.TreeNode

func main() {
	p1 := []int{1, 2, 3}
	p2 := []int{1, 2, 3}
	p := structures.Ints2TreeNode(p1)
	q := structures.Ints2TreeNode(p2)
	// p := &TreeNode{}
	// q := &TreeNode{}
	res := isSameTree(p, q)
	res2 := isSameTree2(p, q)
	fmt.Println("res is", res, res2)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// https://leetcode.cn/problems/same-tree/solutions/2015056/ru-he-ling-huo-yun-yong-di-gui-lai-kan-s-empk
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
