package main

/*
## 변수 초기화
변수 선언은 한 변수 당 하나의 초깃값을 포함할 수 있습니다.
만약 초깃값이 존재한다면, type은 생략될 수 있습니다. 이 경우, 변수는 초깃값의 type을 취합니다.
*/

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}