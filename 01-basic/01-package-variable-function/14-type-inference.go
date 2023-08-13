package main

/*
## Inference type
:= 혹은 var = 표현을 이용해 명시적인 type을 정의하지 않고 변수를 선언할 때, 그 변수 type은 오른 편에 있는 값으로부터 유추됩니다.
변수 선언의 오른 쪽에 무언가 적힐 때, 새로운 변수는 그것과 같은 type이 됩니다.

var i int
j := i // j 는 int
하지만, 오른 편에 type이 정해지지 않은 숫자 상수가 올 때, 새 변수는 그 상수의 정확도에 따라 int 혹은 float64, complex128 이 됩니다.

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
예시에서 v 의 초깃값을 바꿔보시고 v 의 type에 어떤 영향을 끼치는 지 보십시오.
*/

import "fmt"

func main() {
	v := 42;
	fmt.Printf("v is of type %T\n", v)
}