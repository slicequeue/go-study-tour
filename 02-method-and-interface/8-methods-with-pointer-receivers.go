package main

/*
# 값 또는 포인터 리시버 선택
포인터 리시버를 사용하는 데에는 두 가지 이유가 있습니다.

1. 첫번째는, 메서드가 리시버가 가리키는 값을 수정할 수 있기 때문입니다.
2. 두번째는 각각의 메서드 call에서의 value 복사 문제를 피하기 위해서입니다.

예를 들어 리시버가 큰 구조체라면 이것은 더 효율적일 수 있습니다.
이 예시에서, Abs 메서드는 리시버 방식을 수정할 필요가 없지만 Scale 과 Abs 는 모두 *Vertex 라는 리시버 타입으로 되어있습니다.
일반적으로 특정 유형의 모든 방법에는 값이나 포인터 리시버가 있어야 하지만 둘 다 혼합되어서는 안됩니다. (몇 페이지에 걸쳐 그 이유를 알아봅시다.)
*/

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
