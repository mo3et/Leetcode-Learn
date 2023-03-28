//

package leetcode

// import (
// 	"fmt"
// )

func twoSum(num []int, target int) []int {
	m := make(map[int]int)
	for k, v := range num {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}

	return nil
}
