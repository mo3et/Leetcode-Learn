package main

import "fmt"

// F(n)=F(n-1)+F(n-2)
// 先递归，然后记忆化，最后滚动数组

func main() {
	x := climbStairs1(9)
	res := climbStairs2(9)
	fmt.Println(x)
	fmt.Println("ans is", res)
}

// DP 滚动数组实现
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

func climbStairs2(n int) int {
	// 跳过0，直接从 index=2 开始
	d1, d2 := 1, 1
	for i := 2; i <= n; i++ {
		// d2(Fn)= d2(old d2) +d1
		// d1(F(n-1))= d2(old d2)

		d1, d2 = d2, d2+d1
	}
	return d2
}