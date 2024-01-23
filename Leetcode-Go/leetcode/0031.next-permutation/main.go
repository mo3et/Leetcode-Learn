package main

import "fmt"

func main() {
	// nums := []int{1, 1, 5, 6, 4, 7, 2}
	nums := []int{1, 1, 5, 6, 4, 7, 8, 3, 2}
	res := nextPermutation(nums)
	fmt.Println("ans is ", res)
}

// 题解：https://leetcode.cn/problems/next-permutation/solutions/80560/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-

func nextPermutation(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	// 都为从后向前查找
	// 第一个为i,j 查找第一个相邻升序的元素对，满足A[i]<A[j]
	// 再从[j,end]中 查找出满足大于A[i]的元素A[k],进行交换 (此时仍然满足 new i < j)
	// 此时 [j,end] 必然为降序，将其逆置[j,end)，变成升序

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1 // 元素对 (i,j) 比i大的数 k

	// find: A[i]<a[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 { // 不是最后一个排列
		// find: A[i]<A[k]
		for nums[i] >= nums[k] {
			k--
		}
		// swap A[i],A[k]
		nums[i], nums[k] = nums[k], nums[i]
		fmt.Println("i,k swap is ", nums)
	}

	// reverse A[j:end]
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
		fmt.Println("slice reverse is ", nums)
	}
	return nums
}
