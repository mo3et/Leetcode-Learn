package main

import (
	"fmt"

	"github.com/mo3et/leetcode-learn/leetcode-go/structures"
)

type TreeNode = structures.TreeNode

func main() {
	node := &TreeNode{}
	res := preorderTraversal(node)
	fmt.Println("res is", res)
}

func preorderTraversal(root *TreeNode) (res []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}
