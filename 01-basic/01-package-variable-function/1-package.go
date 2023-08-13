package main

/*
# 패키지
모든 Go 프로그램은 패키지로 구성되어있습니다.

프로그램은 main 패키지에서 실행을 시작합니다.
이 프로그램은 import 경로 "fmt" 와 "math/rand" 로 패키지를 사용합니다.
관습적으로, 패키지의 이름은 import 경로의 마지막 요소와 같습니다. 예를 들어 "math/rand" 패키지는 package rand 문으로 시작하는 파일들로 구성되어있습니다.
주의: 이 프로그램들이 실행되는 환경은 결정론적입니다. 따라서 당신이 이 예제 프로그램 실행할 때마다 rand.Intn 는 같은 숫자를 반환할 것입니다.
(다른 숫자를 보기 위해서는 number generator에 seed를 설정하십시오. playground 내에서는 Time은 상수이므로 당신은 seed로서 시간이 아닌 무언가를 사용해야 할 것입니다.
*/

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
