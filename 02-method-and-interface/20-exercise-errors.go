package main
/*
# 실습: Errors
earlier exercise 에서 Sqrt 함수를 복사하고, error 값을 반환하도록 수정합니다.
Sqrt 는 복소수를 지원하지 않으므로 음수가 주어지면 nil이 아닌 오류를 반환해야합니다.
```
type ErrNegativeSqrt float64
```
이라는 새 유형을 만들어, 그것을 error 로 만듭니다.
```
func (e ErrNegativeSqrt) Error() string
```
ErrNegativeSqrt(-2).Error() 는 "cannot Sqrt negative number: -2" 를 반환하는 메소드입니다.
참고: Error 메소드 내에서 fmt.Sprint(e) 를 호출하면 프로그램이 무한 루프를 돌게 됩니다. 
먼저 e 를 fmt.Sprint(float64(e)) 로 변환하여 이를 방지할 수 있습니다. 왜 그럴까요?
음수가 주어지면 ErrNegativeSqrt 값을 반환하도록 Sqrt 함수를 변경합니다.
*/

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	errorValue float64
}

func (e ErrNegativeSqrt) Error() string {
	// fmt.Sprint(e) - 무한루프 오류 빠짐 why? 에러 처리하는 메소드 Error() 내에서 sprint 하기 때문?
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", e.errorValue)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt{x}
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
