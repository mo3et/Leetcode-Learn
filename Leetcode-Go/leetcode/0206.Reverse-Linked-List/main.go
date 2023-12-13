package main

import (
	"github.com/mo3et/leetcode-learn/leetcode-go/structures"
)

type ListNode = structures.ListNode

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	current := reverseList(head)
	// for current != nil {
	// 	fmt.Printf("%d -> ", current.Val)
	// 	current = current.Next
	// }
	// fmt.Println("nil")

	structures.PrintLinkedList(current)
}

func reverseList(head *ListNode) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}

// func printLinkedList(head *ListNode) {
// 	current := head
// 	for current != nil {
// 		fmt.Printf("%d -> ", current.Val)
// 		current = current.Next
// 	}
// 	fmt.Println("nil")
// }
