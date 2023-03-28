package leetcode

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para is Parameter
// one 代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans is answer
// one is first answer
type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {
	qs := []question1{
		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		{
			para1{[]int{3, 2, 4}, 5},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},

		{
			para1{[]int{0, 1}, 1},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 3}, 5},
			ans1{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("-----------------LeetCode Problems ----------------\n")

	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Printf(" [INPUT]: %v 	[OUTPUT]: %v\n", p, twoSum(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
