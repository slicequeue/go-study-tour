package main

/*
# 복수개의 결과
한 함수는 몇 개의 결과든 반환할 수 있습니다.\
swap 함수는 두 개의 string을 반환합니다.
*/

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
