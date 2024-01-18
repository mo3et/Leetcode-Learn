/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import (
	"fmt"

	"github.com/mo3et/leetcode-learn/leetcode-go/structures"
)

type ListNode = structures.ListNode

func main() {
	ln1 := &ListNode{
		Val: 18,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 10,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	res := insertGreatestCommonDivisors(ln1)
	fmt.Println("ans is ", res)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	for cur := head; cur.Next != nil; cur = cur.Next.Next {
		cur.Next = &ListNode{gcd(cur.Val, cur.Next.Val), cur.Next}
	}
	return head
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
