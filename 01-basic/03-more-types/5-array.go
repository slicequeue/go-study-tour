package main

/*
# Arrays (배열)
[n]T 타입은 타입이 T인 n 값들의 배열입니다.
var a [10]int
위 표현은 변수 a를 10개의 정수들의 배열로 선언한 것입니다.
배열의 길이는 그 타입의 일부입니다. 따라서 배열의 크기를 조정할 수 없습니다. 
제약이 많을 것 같지만, 걱정하지 마세요! Go는 배열로 작업할 수 있는 편리한 방법을 제공합니다.
*/

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}
