package main

/*
## 변수
var 문은 변수에 대한 목록을 선언합니다. 함수 인자의 목록이 그랬듯, 마지막은 type입니다.
var 문은 package나 함수 단에 존재할 수 있습니다. 이 예제를 통해 두 경우 모두를 살펴봅시다.
*/

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
