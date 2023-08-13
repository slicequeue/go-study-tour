package main

/*
# Range continued
_ 을 할당하여 인덱스 또는 값을 건너뛸 수 있습니다.

for i, _ := range pow
for _, value := range pow
만약 인덱스만을 원하면, 두 번째 변수를 생략할 수 있습니다.

for i := range pow
*/

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
