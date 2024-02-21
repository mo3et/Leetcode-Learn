package main

import (
	"fmt"

	"github.com/mo3et/leetcode-learn/leetcode-go/structures"
)

type TreeNode = structures.TreeNode

func main() {
	node := &TreeNode{}
	res := postorderTraversal(node)
	fmt.Println("res is", res)
}

func postorderTraversal(root *TreeNode) (res []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
}
