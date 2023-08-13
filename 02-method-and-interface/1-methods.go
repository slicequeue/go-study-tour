package main

/*
# Methods
Go는 클래스를 가지지 않습니다.
하지만, 그와 같은 타입의 메소드를 정의할 수 있습니다.
그 메서드는 특별한 receiver 인자가 있는 함수입니다.
그 receiver는 func 키워드와 메서드 이름 사이의 자체 인수 목록에 나타납니다.
(receiver는 이하 리시버라고 칭하겠습니다.)
이 예에서 Abs 메소드에는 v 라는 이름의 Vertex 유형의 리시버가 있습니다.
*/

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 이 `(v Vertex)` 부분 메소드연결 고리
// v 를 가져다 쓸 수 있다
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 이건 그냥 일반 함수다! 매게 변수는 v Vertex 객체를 받아 활용함
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("v.Abs():",v.Abs())
	fmt.Println("Abs(v):", Abs(v))
}
