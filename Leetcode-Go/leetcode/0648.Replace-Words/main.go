package main

import (
	"fmt"
	"strings"
)

func main() {
	dict := []string{"cat", "bat", "rat"}
	sent := "the cattle was rattled by the battery"
	res := replaceWords(dict, sent)
	fmt.Println("ans is:", res)
}

func replaceWords(dictionary []string, sentence string) string {
	type trie map[rune]trie
	root := trie{}
	for _, s := range dictionary {
		cur := root
		for _, c := range s {
			// fmt.Println("c is:", string(c))
			if cur[c] == nil { // 已经循环完当前单词
				cur[c] = trie{}
				fmt.Println("nil c is:", cur[c])
				// for key := range cur {
				// 	fmt.Println("key is:", string(key))
				// }
			}
			cur = cur[c]
		}
		cur['#'] = trie{}
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		cur := root
		for j, c := range word {
			if cur['#'] != nil { // 到达截取位置
				words[i] = word[:j]
				fmt.Println("words is:", word)
				break
			}
			if cur[c] == nil {
				break
			}
			cur = cur[c]

		}

	}
	return strings.Join(words, " ")
}
