package main

import "fmt"

/*
## 숫자형 상수
숫자형 상수는 매우 정확한 값 입니다.

type이 정해지지 않은 상수는 그것의 문맥에서 필요한 type을 취합니다.

needInt(Big) 을 출력해보세요.

(int 는 최대 64-bit 혹은 때때로 더 작은 정수를 저장할 수 있습니다.)
*/

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
