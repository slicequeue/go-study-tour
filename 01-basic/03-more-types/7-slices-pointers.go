package main

/*
# Slices are like references to arrays (배열을 참조하는 슬라이스)
슬라이스는 어떤 데이터도 저장할 수 없습니다. 이것은 단지 기본 배열의 한 영역을 나타낼 뿐입니다.
슬라이스의 요소를 변경하면 기본 배열의 해당 요소가 수정됩니다.
동일한 기본 배열을 공유하는 다른 슬라이스는 이러한 변경사항을 볼 수 있습니다.
*/

import "fmt"

func main() {
	names := [4]string {
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	b[3] = "!!!" // error
	fmt.Println(names)
	// panic: runtime error: index out of range [3] with length 2
}