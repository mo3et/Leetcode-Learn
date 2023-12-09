package leetcode

import (
	"github.com/mo3et/leetcode-learn/leetcode-go/structures"
)

type ListNode = structures.ListNode

// func main() {
// 	head := &ListNode{
// 		Val: 1,
// 		Next: &ListNode{
// 			Val: 2,
// 			Next: &ListNode{
// 				Val:  3,
// 				Next: nil,
// 			},
// 		},
// 	}
// 	res := reverseList(head)

// 	resV := structures.List2Ints(reverseList(structures.Ints2List(resS)))
// 	resV := make([]int, 0)
// 	for i := 0; i < len(&res); i++ {
// 		resV
// 	}

// 	fmt.Println("ans is ", resV)
// }

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
