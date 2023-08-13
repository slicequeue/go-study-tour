package main

/*
# 이어서 Function
두 개 이상의 연속된 이름이 주어진 함수 매개변수가 같은 type일 때 당신은 마지막 매개변수를 제외한 매개변수들의 type을 생략할 수 있습니다.
이번 예제에서는 우리는 다음을
x int, y int
이렇게 줄였습니다.
x, y int
*/

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
