package main

import "fmt"

// link: https://leetcode.cn/problems/number-of-burgers-with-no-waste-of-ingredients/description/
func main() {
	tomato := 16
	chess := 7
	res := numOfBurgers(tomato, chess)
	fmt.Println("ans is ", res)
}

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 != 0 || tomatoSlices < cheeseSlices*2 || cheeseSlices < tomatoSlices/4 {
		return nil
	}
	biggercnt := tomatoSlices/2 - cheeseSlices
	smallercnt := cheeseSlices - biggercnt
	return []int{biggercnt, smallercnt}
}
