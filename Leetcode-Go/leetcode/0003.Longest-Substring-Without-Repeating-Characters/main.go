package main

import "fmt"

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println("ans is ", res)
}

// 题解：https://leetcode.cn/problems/longest-substring-without-repeating-characters/solutions/227999/wu-zhong-fu-zi-fu-de-zui-chang-zi-chuan-by-leetc-2/

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	// 用 一个 map 来记录字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针rk，从字符串的左边界左侧-1，开始移动(相当于还没开始)
	rk, ans := -1, 0
	// 枚举左指针
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针移动一格，删除一个字符
			delete(m, s[i-1])
		}
		// 移动右指针
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}
