package main

/*
# 타입 스위치
_type_switch는 여러 타입의 선언을 직렬로 허용하는 구조입니다.
타입 스위치는 일반 스위치문과 같지만 타입 스위치문의 경우에는 값이 아닌 타입을 명시하며, 
그 값들은 지정된 인터페이스 값에 의해 유지되는 값의 타입과 비교됩니다.
```
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```
타입 스위치의 선언은 타입 선언 i.(T) 와 같은 구문을 가집니다. 
그러나 특정 타입인 T 는 type 이라는 키워드로 대체됩니다.
이 스위치 문장은 인터페이스 값 i 가 T 형인지 S 형인지 시험합니다. 
T 와 S 의 각각의 경우 변수 v 는 각각 T 형과 S 형으로, i 형이 보유한 값을 보유하게 됩니다. 
디폴트 케이스(일치하지 않는 경우)에서 변수 v 는 인터페이스 종류와 값이 i 와 같습니다.

---
활용도가 높을 것으로 보임! java, js 에서 instanceof  사용했던 것 이런 형태로 활용이 가능하구나!
inserface{} 타입을 주목하자! - 다 받을 수 있으며 이 부분에서 type 을 활용할 수 있으니 말이다!
*/

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}