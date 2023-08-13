package main

/*
# defer 스택 쌓기
연기된 함수 호출들은 스택에 쌓입니다. 
한 함수가 종료될 때 그것의 연기된 함수들은 후입선출 순서로 수행됩니다.
defer문에 대해 더 알고 싶다면 이 [블로그 글](https://go.dev/blog/defer-panic-and-recover)을 읽어보십시오.
---
스택으로 쌓이는 구나...

*/

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}