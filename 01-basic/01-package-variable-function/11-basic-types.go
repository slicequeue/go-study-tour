package main

/**
# 기본 자료형
고의 기본 type들은 다음과 같습니다.
```
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // uint8의 별칭
rune // int32의 별칭
     // 유니코드에서 code point를 의미합니다.
float32 float64
complex64 complex128
```
이 예시는 몇 가지 type의 변수를 보여줄 것이고, 그 변수 선언들은 import 문과 마찬가지로 조각으로 쪼개질 수 있습니다
int 와 uintptr type은 보통 32-bit 시스템에서는 32 bit, 64-bit 시스템에서는 64 bit의 길이입니다. 정수 값이 필요할 때에는 특정한 이유로 크기를 정해야하거나 unsigned 정수 type을 사용해야하는 게 아니라면 int 를 사용해야합니다.
*/

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
