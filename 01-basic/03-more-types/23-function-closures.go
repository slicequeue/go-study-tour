package main

/*
# Function closures (함수 클로저)
Go 함수들은 클로저일 수도 있습니다. 
클로저는 함수의 외부로부터 오는 변수를 참조하는 함수 값입니다. 
함수는 참조된 변수에 접근하여 할당될 수 있습니다. 
이러한 의미에서 함수는 변수에 "bound(바운드)" 됩니다. 
그 예로, adder 함수는 클로저를 반환합니다. 
각 클로저는 그 자체의 sum 변수에 bound(바운드) 되어 있습니다.
*/

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
