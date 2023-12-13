package main

import (
	"fmt"

	"github.com/mo3et/leetcode-learn/leetcode-go/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structures.ListNode

func main() {
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
}

func printLinkedList(head *ListNode) {
	curent := head
	for curent != nil {
		fmt.Printf("%d ->", curent.Val)
		curent = curent.Next
	}
	fmt.Println("nil")
}
