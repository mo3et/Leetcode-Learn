package main

import "fmt"

func main() {
	x := climbStairs1(9)
	fmt.Println(x)
}

func climbStairs1(n int) int {
	q, p, r := 0, 0, 1
	for n > 0 {
		q = p
		p = r
		r = q + p
		n--
	}
	return r
}
