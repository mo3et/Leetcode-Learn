package main

import (
	"fmt"

	"github.com/mo3et/leetcode-learn/leetcode-go/structures"
)

type TreeNode = structures.TreeNode

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	res := diameterOfBinaryTree(node)
	fmt.Println("ans is", res)
}

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		lLen := dfs(node.Left) + 1
		rLen := dfs(node.Right) + 1
		ans = max(ans, lLen+rLen)
		return max(lLen, rLen)
	}
	dfs(root)
	return
}
