package main

/*
# 
Fibonacci closure (피보나치 클로저) 연습하기
한 재밌는 함수를 알아봅시다.
fibonacci 함수를 구현합니다. 
이것은 연속적인 피보나치 수(0, 1, 1, 2, 3, 5, ...)를 반환하는 함수(클로저) 입니다.
*/

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	idx := 0
	return func() int {
		var result int;
		if idx < 2 {
			if idx == 0 { result = a }
			if idx == 1 { result = b }
			idx += 1
		} else {
			result = a + b
			a = b
			b = result
		}
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
