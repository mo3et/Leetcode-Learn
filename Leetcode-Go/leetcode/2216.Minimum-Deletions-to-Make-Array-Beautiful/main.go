package main

import "fmt"

// 还可以用 栈模拟和贪心
func main() {
	// nums := []int{1, 1, 2, 3, 5}
	nums := []int{1, 1, 2, 2, 3, 3}
	res := minDeletion(nums)
	fmt.Println("ans is", res)
}

// 用了栈来解决
func minDeletion(nums []int) (res int) {
	n := len(nums)
	if n == 0 {
		return 0
	}
	st := make([]int, 0)
	st = append(st, nums[0])
	for i, index := 1, 0; i < n; i++ {
		// if index%2==0&& nums[i]!=a[j]{
		//         a=append(a,nums[i])
		//         j++
		// }
		fmt.Println(st)
		fmt.Println("index is ", index)
		if index%2 == 0 && nums[i] == st[index] {
			res++
		} else {
			st = append(st, nums[i])
			index++
		}
	}
	if len(st)%2 != 0 {
		res++
	}

	return res
}
