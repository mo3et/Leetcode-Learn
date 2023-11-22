package main

import "strconv"

func countSeniors(details []string) int {
	count := 0
	// for i := 0; i < len(details); i++ {
	for _, detail := range details {
		// age, _ := strconv.Atoi(detail[i][11:13])
		age, _ := strconv.Atoi(detail[11:13])
		if age > 60 {
			count++
		}
	}
	return count
}

// TODO 如何使用 slice 然后把一组数据append进去 例如有三个以上的 slice内容
func main() {
	details := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}
	countSeniors(details)
}
