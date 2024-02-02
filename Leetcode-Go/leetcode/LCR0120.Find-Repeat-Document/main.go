package main

import "fmt"

// Link: https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

func main() {
	documents := []int{2, 2, 3, 0, 5, 0, 5}
	res := findRepeatDOcument(documents)
	fmt.Println("ans is", res)
}

func findRepeatDOcument(documents []int) int {
	hash := make([]int, len(documents))
	for _, n := range documents {
		// fmt.Println(hash)
		if hash[n] > 0 {
			return n
		} else {
			hash[n] = 1
		}
	}
	return -1
}
