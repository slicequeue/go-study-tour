package main

/*
이어서 For
초기화 구문과 사후 구문은 필수는 아닙니다.
*/

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}