package main

/*
# 메소드와 포인터 indirection (2)
이와 동등한 일은 역방향에서 일어날 수 있다.
값 인수를 사용하는 함수는 다음과 같은 특정 유형의 값을 사용해야 합니다.
```
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```
value receiver 가 있는 메서드는 다음과 같이 호출될 때, 
값이나 포인터를 리시버로 사용합니다.
```
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```
이 경우에, p.Abs() 라는 메서드는 (*p).Abs() 로 해석됩니다.
*/

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("v.Abs():", v.Abs())
	fmt.Println("AbsFunc(v):", AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println("p.Abs():", p.Abs())
	fmt.Println("AbsFunc(*p):", AbsFunc(*p))
}
