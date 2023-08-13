package main

/*
#Import
오른쪽의 코드는 괄호로 import 를 그룹 짓습니다. 이를 "factored" import 문이라고합니다.
당신은 여러 import문들을 다음과 같이 적을 수도 있습니다.
import "fmt"
import "math"
하지만 위의 방식보다는 오른쪽처럼 "factored" import문을 사용하는 것이 좋은 스타일입니다.
*/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}