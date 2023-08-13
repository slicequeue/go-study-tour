package main

/*
# Structs (구조체)
구조체는 필드의 집합체입니다.
*/

import "fmt"

type Vertex struct {
	X int
	Y int
	// z *
}

func main() {
	fmt.Println(Vertex{1, 2})
}
