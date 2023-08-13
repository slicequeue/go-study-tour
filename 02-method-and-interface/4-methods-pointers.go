package main

/*
# 포인터 리시버
포인터 리시버로 메소드를 선언 할 수 있습니다.
이는 리시버 유형이 일부 유형 T 에 대한 리터럴 구문 *T 를 가짐을 의미합니다. 
(또한 T 자체는 *int 와 같은 포인터가 될 수 없습니다.)
예를 들어, 여기서 Scale 메소드는 *Vertex 에 정의되어 있습니다.
포인터 리시버가 있는 메소드는 Scale 처럼 리시버가 가리키는 값을 수정할 수 있습니다.
메소드는 종종 리시버를 수정해야하기에 포인터 리시버가 값 리시버보다 더 일반적입니다.
16 행의 Scale 함수 선언에서 * 를 제거하고 프로그램의 동작이 어떻게 변하는 지 관찰해보십시오.
값 수신기를 사용하면 Scale 메서드가 원래 Vertex 값의 복사본에서 작동합니다. 
(이것은 다른 함수 인수와 동일합니다.)
Scale 메소드에는 main 함수에 선언된 Vertex 값을 변경하기 위한 포인터 리시버가 있어야합니다.
*/

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	v.Scale(10) // 이렇게 한다고 해서 값이 반영되지는 않는다
	fmt.Println(v.Abs())
	fmt.Println(v)
}