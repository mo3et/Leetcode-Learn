package main

import "fmt"

// 题解: https://leetcode.cn/problems/beautiful-towers-i/solutions/2614597/mei-li-ta-i-by-leetcode-solution-uqnf

func main() {
	maxHeights := []int{5, 3, 4, 1, 1}
	res := maximunSumOfHeights(maxHeights)
	fmt.Println("ans is:", res)
}

func maximunSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	prefix := make([]int, n)
	suffix := make([]int, n)
	stack1, stack2 := []int{}, []int{}

	for i := 0; i < n; i++ {
		for len(stack1) > 0 && maxHeights[i] < maxHeights[stack1[len(stack1)-1]] {
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack1) == 0 {
			prefix[i] = (i + 1) * maxHeights[i]
		} else {
			last := stack1[len(stack1)-1]
			prefix[i] = prefix[last] + (i-last)*maxHeights[i]
		}
		stack1 = append(stack1, i)
	}

	res := 0
}
