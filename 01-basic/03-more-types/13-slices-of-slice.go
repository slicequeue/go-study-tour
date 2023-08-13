package main

/*
Slices of slices (슬라이스의 슬라이스)
슬라이스는 다른 슬라이스를 포함하여 모든 타입을 담을 수 있습니다.
*/

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	
}
