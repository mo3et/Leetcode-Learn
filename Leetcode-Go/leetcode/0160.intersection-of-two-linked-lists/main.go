package main

import (
	"fmt"

	"github.com/mo3et/leetcode-learn/leetcode-go/structures"
)

type ListNode = structures.ListNode

// link https://leetcode.cn/problems/intersection-of-two-linked-lists/description/
func main() {
	// ！！ints2List 有问题
	headAori := []int{4, 1, 8, 4, 5}
	headA := structures.Ints2List(headAori)
	headBori := []int{5, 6, 1, 8, 4, 5}
	headB := structures.Ints2List(headBori)
	res := getIntersectionNode(headA, headB)
	res2 := getIntersectionNodeDouPointer(headA, headB)
	fmt.Println("ans is", res)
	fmt.Println("ans is", res2)
}

// Double Pointer
func getIntersectionNodeDouPointer(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pointA, pointB := headA, headB
	for pointA != pointB {
		if pointA == nil {
			pointA = headB
		} else {
			pointA = pointA.Next
		}
		if pointB == nil {
			pointB = headA
		} else {
			pointB = pointB.Next
		}
	}
	return pointA
}

// Hash Set
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 遍历 headA 并存储在map
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}

	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}
