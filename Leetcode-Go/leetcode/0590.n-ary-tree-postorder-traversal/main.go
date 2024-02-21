package main

import "fmt"

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	node := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{Val: 5, Children: nil},
					{Val: 6, Children: nil},
				},
			},
			{
				Val:      2,
				Children: nil,
				// Children: []*Node{
				// 	{Val: 7, Children: nil},
				// },
			},
			{
				Val:      4,
				Children: nil,
			},
		},
	}
	res := postorder(node)
	fmt.Println("ans is", res)
}

func postorder(root *Node) (ans []int) {
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, ch := range node.Children {
			dfs(ch)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return
}
