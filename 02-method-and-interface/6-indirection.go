package main

/*
# 메소드와 포인터 indirection
이전 두 프로그램을 비교하여 포인터 인수의 함수가 다음과 같은 포인터를 사용해야함을 알 수 있습니다.
```
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```
포인터 리시버가 있는 메소드는 다음과 같이 호출될 때 값이나 포인터를 리시버로 받아들입니다.
```
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```
v 라는 문장의 경우, v.Scale(5) 는 v 가 포인터가 아니라 값인데도 포인터 리시버가 있는 메서드는 자동으로 호출됩니다.
즉, Scale 메서드가 포인터 리시버를 가졌기 때문에 편의상 Go는 v.Scale(5) 라는 것을 (&v).Scale(5) 로 해석합니다.
*/

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}