package main

/*
# 
맵 연습하기
WordCount 를 구현해봅시다. 
이것은 문자열 s 의 각 word 의 개수의 맵을 반환해야 합니다. 
wc.Test 함수는 제공된 함수를 테스트하고 성공 혹은 실패를 출력하는 함수입니다. 
strings.Fields가 도움이 될 수 있습니다.
*/

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
