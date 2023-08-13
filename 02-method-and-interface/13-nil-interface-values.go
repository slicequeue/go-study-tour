package main

/*
# Nil 인터페이스 값
Nil 인터페이스 값은 값 또는 콘크리트 유형 모두를 가지지 않습니다.
nil 인터페이스에서 메소드를 호출하는 것은 런타임 에러입니다. 
왜냐하면, 어떠한 구체적인 메소드를 호출할지를 나타내는 인터페이스 튜플 내부의 타입이 없기 때문입니다.
*/

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
