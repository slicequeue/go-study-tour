package main

/*
인터페이스의 암시적 구현
type implements는 메소드를 실행함으로써 인터페이스를 구현합니다. 
명시적 intent의 선언도, "implementation"의 키워드도 없습니다.
암시적 인터페이스는 인터페이스의 정의를 구현으로부터 분리하며, 
이는 사전 정렬 없이 어떠한 패키지에 등장할 수 있습니다.
*/

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
