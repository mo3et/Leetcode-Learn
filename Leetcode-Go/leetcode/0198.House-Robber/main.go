package main

func main() {
}

func rob(nums []int) int {
	d0, d1 := 0, 0
	for _, v := range nums {
		d := max(d0, d1+v)
		d0 = d1
		d1 = d
	}
	return d1
}
