package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "&amp; is an HTML entity but &ambassador; is not."
	// text1 := "and I quote: &quot;...&quot;"
	res := entityParser(text)
	fmt.Println("ans is ", res)
}

func entityParser(text string) string {
	// 定义需要处理的字符串内容为map
	entityMap := map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&amp;":   "&",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
	}
	idx := 0
	n := len(text)
	// 保存 Parse 后的 String 内容
	res := make([]string, 0)

	// 遍历的text index 到 len
	for idx < n {
		// 开始时设置默认idx Val不为实体
		// 以满足后面不为 Entity 的condition
		isEntity := false
		if text[idx] == '&' {
			// 在 Map 中找出对应的实体内容
			for k, v := range entityMap {
				// 满足index+len(k) <=n (不超过text总长度，不会越界)
				// 并满足 text[index:index+len(k)]符合实体内容
				if idx+len(k) <= n && text[idx:idx+len(k)] == k {
					res = append(res, v)
					// isEntity为T，本次遍历不会再执行下面的 !isEntity
					isEntity = true
					idx += len(k)
					break
				}
			}
		}
		// 如果不为实体，就把当前字符加入新Slice，index 进入下一位
		if !isEntity {
			res = append(res, text[idx:idx+1])
			idx++
		}
	}
	fmt.Println("res is ", res)

	// 将Slice 拼接成 string 返回
	return strings.Join(res, "")
}
