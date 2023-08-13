package main

/*
# Mutating Maps
m 맵에 요소를 추가하거나 업데이트하기:
```
m[key] = elem
```
요소 검색하기:
```
elem = m[key]
```
요소 제거하기:
```
delete(m, key)
```
두 개의 값을 할당하여 키가 존재하는지 테스트할 수 있습니다.
```
elem, ok = m[key]
```
만약 key 가 m 안에 있다면, ok 는 true 입니다. 아니라면, ok 는 false 입니다.
만약 key 가 맵 안에 없다면, elem 은 map의 요소 타입의 zero value입니다.
참고: 만약 elem 또는 ok 가 아직 선언되지 않았다면, 아래와 같은 짧은 정의식을 사용할 수 있습니다.
*/

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
