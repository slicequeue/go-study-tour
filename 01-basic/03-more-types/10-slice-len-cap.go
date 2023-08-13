package main

/*
# Slice length and capacity (슬라이스의 길이와 용량)
슬라이스는 _length(길이)_와 _capacity(용량)_를 둘다 가지고 있습니다.
슬라이스의 길이란 슬라이스가 포함하는 요소의 개수입니다.
슬라이스의 용량이란 슬라이스의 첫 번째 요소부터 계산하는 기본 배열의 요소의 개수입니다.
슬라이스 s 의 길이와 용량은 len(s) 와 cap(s) 식으로 얻을 수 있습니다.
슬라이스의 길이를 연장하기 위해서는 슬라이스에 충분한 용량이 있을 때, 다시 슬라이싱을 하면 됩니다.
예제의 슬라이스 작업중 하나를 슬라이스의 용량을 넘어가도록 변경하여 어떤 일이 발생하는지 확인해 보세요.
*/

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
