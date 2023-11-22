package main

import "fmt"

func main() {
	rc := []int{18, 43, 36, 13, 7}
	a1 := maximunSum(rc)
	fmt.Println(a1)
}

func maximunSum(nums []int) int {
	ans := -1
	record := make([]int, 82)
	for _, num := range nums {
		s := 0
		for x := num; x > 0; x /= 10 {
			s += x % 10
		}
		if record[s] > 0 {
			ans = max(ans, record[s]+num)
		}
		record[s] = max(record[s], num)
		fmt.Print(record[s], " ")
	}
	return ans
}
