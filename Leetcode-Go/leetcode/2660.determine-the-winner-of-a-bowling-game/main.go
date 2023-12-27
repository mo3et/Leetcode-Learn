package main

import "fmt"

func main() {
	player1 := []int{4, 10, 7, 9}
	player2 := []int{6, 5, 2, 3}
	res := isWinner(player1, player2)
	fmt.Println("ans is ", res)
}

func isWinner(player1, player2 []int) int {
	p1, p2 := score(player1), score(player2)
	if p1 > p2 {
		return 1
	} else if p1 < p2 {
		return 2
	} else {
		return 0
	}
}

func score(player []int) int {
	res := 0
	for i, x := range player {
		if i > 0 && player[i-1] == 10 || i > 1 && player[i-2] == 10 {
			res += 2 * x
		} else {
			res += x
		}
	}
	return res
}
