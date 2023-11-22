package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	x1 := rob(nums)
	fmt.Println(x1)
}

func rob(nums []int) int {
	d0, d1 := 0, 0
	for _, v := range nums {
		d := max(d0+v, d1)
		d0 = d1
		d1 = d
	}
	return d1
}
