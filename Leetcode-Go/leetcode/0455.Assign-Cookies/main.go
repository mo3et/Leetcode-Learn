package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2, 3}
	g2 := []int{10, 9, 8, 7}
	s := []int{1, 1}
	s2 := []int{5, 6, 7, 8}
	res := findContentChildren(g, s)
	res2 := findContentChildrenSimp(g2, s2)
	fmt.Println("ans is ", res)
	fmt.Println("ans2 is ", res2)
}

// g is childrens, s is cookies
func findContentChildren(g []int, s []int) (res int) {
	sort.Ints(g)
	sort.Ints(s)
	m, n := len(g), len(s)
	// i 为children的index，j 为cookie的index
	for i, j := 0, 0; i < m && j < n; i++ {
		// 排除当前 cookie index小于len,且children需要的值大于cookie的值
		for j < n && g[i] > s[j] {
			// 跳到下一个cookie
			j++
		}
		// ~~ cookie 满足 children的需求~~
		// 只需判断 index 是否在len内，就必然满足<=cookie Val
		if j < n {
			res++
			j++
		}
	}
	return
}

func findContentChildrenSimp(g []int, s []int) (res int) {
	sort.Ints(g)
	sort.Ints(s)
	m, n := len(g), len(s)
	// i 为children的index，j 为cookie的index
	for i, j := 0, 0; i < m && j < n; {
		if g[i] <= s[j] {
			res++
			i++
		}
		j++
	}
	return
}
