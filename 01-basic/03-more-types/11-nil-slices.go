package main

/*
# Nil slices (nil 슬라이스)
슬라이스의 zero value는 nil 입니다.
nil 슬라이스의 길이와 용량은 0이며, 기본 배열을 가지고 있지 않습니다.
*/

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
