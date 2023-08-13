package main

/*
# Pointers (포인터)
Go는 포인터를 지원합니다. 포인터는 값의 메모리 주소를 가지고 있습니다.

*T 타입은 T 값을 가리키는 포인터입니다. 이것의 zero value는 nil 입니다.

var p *int
& 연산자는 이것의 피연산자에 대한 포인터를 생성합니다.

i := 42
p = &i
* 연산자는 포인터가 가리키는 주소의 값을 나타냅니다.

fmt.Println(*p) // 포인터 p를 통해 i 읽기
*p = 21         // 포인터 p를 통해 i 설정
이것은 "역 참조" 또는 "간접 참조"로 알려져 있습니다.

C언어와는 다르게, Go는 포인터 산술을 지원하지 않습니다.
*/

import "fmt"

func main() {
	i, j := 42, 2701
	
	p := &i // point to i
	fmt.Println(p)
	fmt.Println(*p) // read i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // pint to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

}
