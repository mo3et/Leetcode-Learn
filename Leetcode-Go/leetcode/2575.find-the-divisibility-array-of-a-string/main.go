package main

import "fmt"

func main() {
	word := "998244353"
	m := 3
	result := divisibilityArray(word, m)
	fmt.Println("ans is", result)
}

// TODO 还是不太懂

func divisibilityArray(word string, m int) []int {
	ans := make([]int, len(word))
	x := 0
	// c1 := 0
	for i, char := range word {
		// fmt.Println("x is", x)
		// c1 = (x*10 + int(char-'0'))
		// fmt.Println("changed x is", c1)
		x = (x*10 + int(char-'0')) % m
		if x == 0 {
			ans[i] = 1
		}
	}
	return ans
}
