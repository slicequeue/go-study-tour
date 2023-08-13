package main

/*
# 포인터와 함수
여기서 우리는 함수로 재 작성된 Abs 와 Scale 메소드를 볼 수 있습니다.
다시 16 행에서 * 를 제거해보십시오.
동작이 바뀌는 것을 볼 수 있습니까?
예제를 컴파일하기 위해서는 무엇을 변경해야합니까?
(확실하지 않은 경우 다음 페이지로 이동하십시오.)
*/

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	Scale(v, 10)
	fmt.Println(Abs(v))
	fmt.Println(v)
}
