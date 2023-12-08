package main

import "fmt"

func main() {
	s := "abcdba"
	res := validPalindrome(s)
	fmt.Println("ans is ", res)
}

func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			// 删除r位置字符后，是否回文
			for l, r := low, high-1; l < r; l, r = l+1, r-1 {
				if s[l] != s[r] {
					flag1 = false
					break
				}
			}
			// 删除l位置字符后，是否回文
			for l, r := low+1, high; l < r; l, r = l+1, r-1 {
				if s[l] != s[r] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
