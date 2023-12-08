package main

import "fmt"

func main() {
	b1 := []int{5, 5, 5, 5, 10, 20}
	// b := []int{5, 5, 10, 10, 20}
	res := lemonadeChange(b1)
	fmt.Println("ans is", res)
}

// map 解法，还可以直接用 five,ten 直接计数
func lemonadeChange(bills []int) bool {
	cash := make(map[int]int)
	for _, b := range bills {
		if b == 5 {
			cash[b]++
		}
		if b == 10 {
			if cash[5] > 0 {
				cash[5]--
				cash[10]++
			} else {
				return false
			}
		}
		if b == 20 {
			if cash[10] > 0 && cash[5] > 0 {
				cash[10]--
				cash[5]--
				cash[20]++
			} else if cash[10] == 0 && cash[5] > 2 {
				cash[5] -= 3
			} else {
				return false
			}
		}
	}
	fmt.Println(cash)
	return true
}
