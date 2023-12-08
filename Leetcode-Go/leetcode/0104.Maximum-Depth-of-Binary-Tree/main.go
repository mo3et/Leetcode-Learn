package main

import (
	"fmt"

	"github.com/mo3et/leetcode-learn/leetcode-go/structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structures.TreeNode

func main() {
	root := &TreeNode{} //{Val: {3, 9, 20, nil, nil, 15, 7}}
	x1 := maxDepth(root)
	fmt.Println(x1)
}

func maxDepth(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, cnt int) {
		if node == nil {
			return
		}
		cnt++
		ans = max(ans, cnt)
		dfs(root.Left, cnt)
		dfs(root.Right, cnt)
	}
	dfs(root, 0)
	return
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)
	return max(lDepth, rDepth) + 1
}
