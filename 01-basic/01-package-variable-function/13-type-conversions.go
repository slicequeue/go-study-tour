package main

/*
## Type 변환
T(v) 는 v 라는 값을 T type으로 변환시켜줍니다.

몇 가지 숫자 변환 예시:

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
혹은 좀 더 간단히:

i := 42
f := float64(i)
u := uint(f)
C와 달리 Go는 다른 type의 요소들 간의 할당에는 명시적인 변환을 필요로합니다. 예시에서 float64 나 uint 변환을 제거해보시고, 어떤 일이 발생하는 지 보십시오.
- .\type-conversions.go:28:15: cannot use f (variable of type float64) as uint value in variable declaration
*/

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}