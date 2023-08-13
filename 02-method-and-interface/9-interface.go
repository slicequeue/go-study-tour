package main

/*
# Interfaces
_interface_type_은 메소드의 시그니처 집합으로 정의됩니다.
인터페이스 유형의 값은 해당 메소드를 구현하는 모든 값을 보유 할 수 있습니다.
참고: 예제 코드 22행에 오류가 있습니다. Vertex (값 유형)는 Abser 를 구현하지 않습니다. 
	왜냐하면 Abs 메소드는 *Vertex (포인터 유형)에서만 정의되기 때문입니다.
*/

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v // 이 부분 주목! 실제 값타입에서는 오류가 발생함 struct 타입이기 때문?

	fmt.Println(a.Abs())
}
