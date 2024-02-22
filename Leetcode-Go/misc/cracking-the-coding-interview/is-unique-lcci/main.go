package main

import "fmt"

func main() {
	s := "leetcode"
	res := isUnique(s)
	fmt.Println("ans is", res)
}

func isUnique(astr string) bool {
	store := [26]int{}
	for _, s := range astr {
		if store[s-'a'] > 0 {
			return false
		}
		store[s-'a']++
	}
	return true
}
