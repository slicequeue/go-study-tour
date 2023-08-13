package main

/*
# 타입 선언
_type_assertion_은 인터페이스 값의 기초적인 콘크리트 값에 대한 접근을 제공합니다.
```
t := i.(T)
```
이는 인터페이스 값 i 가 콘크리트 타입 T 를 갖고 있으며, 그 기본 값인 T 값을 변수 t 에 할당하고 있다고 선언합니다.
만약 i 가 T 를 갖지 못하면면 그 선언은 panic 상태가 됩니다.
인터페이스 값이 특정 유형을 보유하는지 여부를 _test_하기 위해, 타입 선언에서는 두 가지 값, 즉 기본 값과 선언 성공 여부를 보고하는 부울 값을 반환할 수 있습니다.
```
t, ok := i.(T)
```
만약 i 가 T 를 갖고 있다면, t 는 underlying value가 되며, ok 가 true를 반환합니다.
만약 그렇지 않다면, ok 는 거짓이 되고 t 는 T 라는 유형의 zero 값이 되며 어떠한 패닉도 발생하지 않습니다.
이 구문과 map에서 읽는 구문 간의 유사성에 유의하십시오.

---

이걸 언제쓰는거지?

*/

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
