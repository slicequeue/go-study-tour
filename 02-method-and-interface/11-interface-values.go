package main

/*
# 인터페이스 값
hood 아래에서, 인터페이스의 값은 값과 콘크리트 타입의 튜플이라고 생각할 수 있습니다.
```
(value, type)
```
인터페이스 값은 특정 기초 콘크리트 유형의 가치를 가집니다.
인터페이스 값으로 메소드를 호출하면 기본 형식에 동일한 이름의 메서드가 실행됩니다.
*/

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}


func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}