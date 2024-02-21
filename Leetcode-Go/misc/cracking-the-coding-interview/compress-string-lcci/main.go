package main

import (
	"fmt"
	"strconv"
	"strings"
)

// link: https://leetcode.cn/problems/compress-string-lcci/description/?envType=study-plan-v2&envId=cracking-the-coding-interview

func main() {
	S := "aabcccccaaa"
	S1 := "aabbc"
	res := compressString(S)
	res2 := compressString(S1)
	fmt.Println("ans is", res, "\n", res2)
}

func compressString(S string) string {
	if len(S) == 0 {
		return S
	}

	var builder strings.Builder
	count := 1

	// 从第二索引(i=1)开始计算,所以是小于len(S),而不-1
	// count为1是避免当 第一个字符不同时，count为0 没有累加
	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			count++
		} else {
			builder.WriteByte(S[i-1])
			builder.WriteString(strconv.Itoa(count))
			// 拼接完成后重置 count
			count = 1
		}
	}

	// 需要单独处理最后一组相邻重复字符 因为i无法等于len(S)
	builder.WriteByte(S[len(S)-1])
	builder.WriteString(strconv.Itoa(count))

	/* 不需要单独处理的写法 */
	/* 当 i < len(S) 就会累加, 最后一位默认处理  */
	// for i := 1; i <= len(S); i++ {
	// 	if i < len(S) && S[i] == S[i-1] {
	// 		count++
	// 	} else {
	// 		builder.WriteByte(S[i-1])
	// 		builder.WriteString(strconv.Itoa(count))
	// 		count = 1
	// 	}
	// }

	if builder.Len() < len(S) {
		return builder.String()
	}
	return S
}
