// 2697. Lexicographically Smallest Palindrome

// You are given a string s consisting of lowercase English letters, and you are allowed to perform operations on it. In one operation, you can replace a character in s with another lowercase English letter.
// Your task is to make s a palindrome with the minimum number of operations possible. If there are multiple palindromes that can be made using the minimum number of operations, make the lexicographically smallest one.
// A string a is lexicographically smaller than a string b (of the same length) if in the first position where a and b differ, string a has a letter that appears earlier in the alphabet than the corresponding letter in b.
// Return the resulting palindrome string.

package main

import "fmt"

func main() {
	s := "abcabcbb"
	res := makeSmallestPalindrome(s)
	fmt.Println("ans is ", res)
}

func makeSmallestPalindrome(s string) string {
	// 判断和修改回文字 只需要从端点同时出发对比即可
	left, right := 0, len(s)-1
	// 用来修改的slice
	t := []byte(s)
	for left < right {
		if s[left] != s[right] {
			t[left] = min(s[left], s[right])
			t[right] = t[left]
		}
		left++
		right--
	}
	return string(t)
}
