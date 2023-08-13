package main

/*
# 빈 인터페이스 값
0 메서드를 지칭하는 인터페이스 유형을 _empty_interface_라고 합니다:
```
interface{}
```
빈 인터페이스는 모든 유형의 값을 가질 수 있습니다. 
(모든 유형은 최소 0개의 메소드를 구현합니다.)
빈 인터페이스는 알 수 없는 값을 처리하는데 이용됩니다. 
예를 들어, fmt.Print 는 interface{} 타입의 어떠한 인수라도 취할 수 있습니다.
---
"예를 들어, fmt.Print 는 interface{} 타입의 어떠한 인수라도 취할 수 있습니다." - object 같기도 하면서... 다 받을 수 있는 것인가... 
- 특정 활용이 있을 것 처럼 보인다
*/

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}