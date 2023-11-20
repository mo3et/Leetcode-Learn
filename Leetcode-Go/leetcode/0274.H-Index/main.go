package leetcode

import "sort"

func hIndex(citations []int) int {
	minCit, maxCit := len(citations), 0
	for _, v := range citations {
		minCit = min(minCit, v)
		maxCit = max(maxCit, v)
	}
	res := sort.Search(maxCit-minCit+1, func(cit int) bool {
		cit += minCit
		count := 0
		for _, v := range citations {
			if v >= cit {
				count++
			}
		}
		return count < cit
	})
	return res - 1 + minCit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
