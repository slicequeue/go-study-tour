package main

/*
# 짧은 구문의 If
for 과 마찬가지로, if 문 또한 조건문 전에 수행될 짧은 구문으로 시작될 수 있습니다.
그 짧은 구문에서 선언된 변수들은 오직 if 문의 끝까지만 스코프가 존재합니다.
(마지막 return 문에서 v 를 사용해보십시오.)
*/

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
