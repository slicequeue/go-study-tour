package main

/*
# 함수
함수는 0개 혹은 그보다 많이 인자를 받을 수 있습니다.
이번 예시에서 add 는 2개의 int 형 매개변수를 이용합니다.
변수 이름 뒤에 type 이 온다는 것을 명심하세요.
(type이 왜 그런 방식으로 보여지는 지에 대한 자세한 내용은 Go의 선언 syntax에 대한 글에서 확인해보십시오.)
*/

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
