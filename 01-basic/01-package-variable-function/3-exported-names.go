package main

/*
# Export되는 이름
Go에서는 대문자로 시작하는 이름이 export 됩니다. 예를 들어 Pi 가 math 패키지에서 export 되었듯이 Pizza 는 export되는 이름입니다.
pizza 와 pi 는 대문자로 시작하지 않으므로 export되지 않습니다.
pacakge를 import 할 때, 당신은 그 패키지의 export된 이름들만 참조할 수 있습니다. export되지 않은 이름들에는 패키지의 밖에서 접근할 수 없습니다.
코드를 실행해보십시오. 에러메시지를 보십시오.
에러를 고치기 위해서는 math.pi 를 math.Pi 로 변경하시고, 다시 시도해보십시오.
*/

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(math.Pi)
	fmt.Println(math.pi) // .\3-exported-names.go:19:19: undefined: math.pi
}