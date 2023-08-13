package main

/*
# Nil 인터페이스 값
인터페이스 자체 내부의 콘크리트 값이 0일 경우, 그 메소드는 nil 리시버로 호출됩니다.
일부 언어에서는 이것이 null 포인터 예외를 발생시키지만, 
Go 에서는 nil 리시버로 호출되는 것으로 불리는 매우 좋은 방법을 사용하는 것이 일반적입니다. 
(이 예시의 M 이라는 방법과 같습니다.)
nil 콘크리트 값을 갖는 인터페이스 값 자체가 nil이 아니라는 점에 유의해야 합니다.
*/

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		// t 가 nil 일 경우 처리 방법 별도 지정
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
