package main

/*
# Methods (Continued)
구조체가 아닌 형식에 대해서도 메소드를 선언할 수 있습니다.
이 예에서는 Abs 메소드가 있는 숫자 유형 MyFloat 을 확인할 수 있습니다.
메소드와 동일한 패키지에 유형이 정의된 수신자가 있는 메소드만 선언할 수 있습니다.
유형이 다른 패키지 (int 와 같은 빌트인 유형 포함)에 정의된 리시버로 메소드를 선언 할 수 없습니다.
*/

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
