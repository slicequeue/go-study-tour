package main

/*
#  Struct Literals
구조체 리터럴은 필드 값을 나열하여 새로 할당된 구조체 값을 나타냅니다.
Name: 구문으로 필드의 하위 집합만 나열할 수 있습니다. (명명된 필드의 순서는 무관합니다.)
특별한 접두사 & 은 구조체 값으로 포인터를 반환합니다.
*/

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1,2}	// has type
	v2 = Vertex{X:1}	// Y:0 is implict
	v3 = Vertex{}		// X:0 and Y:0
	p = &Vertex{1,2}	// has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
