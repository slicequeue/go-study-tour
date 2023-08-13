package main

/*
## 상수
상수는 변수처럼 선언되지만 const 키워드와 함께 선언됩니다.

상수는 character 혹은 string, boolean, 숫자 값이 될 수 있습니다.

상수는 := 를 통해 선언될 수 없습니다.
*/

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}