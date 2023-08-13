package main

/*
# Appending to a slice (슬라이스에 요소 추가하기)
슬라이스에 새로운 요소를 추가하는 것이 일반적입니다. 
따라서 Go는 내장된 append 함수를 제공합니다. 문서에 append 내장 함수의 설명이 있습니다.
```
func append(s []T, vs ...T) []T
```
append 의 첫번째 파라미터 s 는 슬라이스의 타입 T 입니다. 
그리고 나머지 T 값들은 슬라이스에 추가할 값들입니다.
append 의 결과 값은 원래 슬라이스의 모든 요소와 추가로 제공된 값들을 포함하는 슬라이스입니다.
만약 s 의 원래 배열이 너무 작아서 주어진 값을 모두 추가할 수 없을 경우, 더 큰 배열이 할당됩니다. 
이 때 반환된 슬라이스는 새로 할당된 배열을 가리킵니다.
(슬라이스에 대해 더 알고 싶다면 Slices: usage and internals 글을 읽어보세요! )
*/

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}


func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}