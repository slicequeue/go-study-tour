package main

/*
## 짧은 변수 선언
함수 내에서는 := 라는 짧은 변수 선언은 암시적 type으로 var 선언처럼 사용될 수 있습니다.

함수 밖에서는 모든 선언이 키워드(var, func, 기타 등등)로 시작하므로 := 구문은 사용할 수 없습니다.
*/

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
