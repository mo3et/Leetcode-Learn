//

package leetcode

// package main

import "fmt"

// import (
// 	"fmt"
// )

func twoSum(num []int, target int) []int {
	m := make(map[int]int)  // 0 x1, 1 x2, 2 x3, 3 x4
	for k, v := range num { //k is i, v is num[i]
		fmt.Println(m)
		// fmt.Println(k, v)
		// fmt.Println("Its ", m[target-v], "and m[k]", m[k], " and m[v]", m[v])

		// m[target-v] // == m[another]  another:=target-num[i]
		fmt.Println(m[v])

		if another, ok := m[target-v]; ok { //idx is target - value, or name another
			fmt.Println([]int{k, v})
			fmt.Println([]int{another, k})
			return []int{another, k}

		}
		m[v] = k
	}

	return nil
}

//  ------------ Another Solves --------------

// func twoSum(nums []int, target int) []int {
// 	m := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		another := target - nums[i]
// 		if _, ok := m[another]; ok {
// 			return []int{m[another], i}
// 		}
// 		m[nums[i]] = i
// 	}
// 	return nil
// }

// ============= Solve2 ===================
//  Use range
// func twoSum(nums []int, target int) []int {
// 	m := make(map[int]int)
// 	for k, v := range nums {
// 		if idx, ok := m[target-v]; ok {
// 			return []int{idx, k}
// 		}
// 		m[v] = k
// 	}
// 	return nil
// }

// Go run test (No )

// func main() {

// 	para1 := []int{0, 8, 7, 3, 3, 4, 2}
// 	twoSum(para1, 11)
// 	// ans1{[]int{1, 3}},
// }
