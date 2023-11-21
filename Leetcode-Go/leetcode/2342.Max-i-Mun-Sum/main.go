package main

func main() {
}

func maximunSum(nums []int) int {
	ans := -1
	record := make([]int, 82)
	for _, num := range nums {
		s := 0
		for x := num; x > 0; x /= 10 {
			s += x % 10
		}
		if record[s] > 0 {
			ans = max(ans, record[s]+num)
		}
	}
	return ans
}
