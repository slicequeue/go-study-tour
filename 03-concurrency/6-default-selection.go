package main

/*
# Default Selection
select 에서의 default case 는 다른 case들이 모두 준비되지 않았을 때 실행됩니다.
block 없이 전송이나 수신을 수행하도록 default case를 사용해봅시다.
```
select {
case i := <-c:
    // use i
default:
    // c로부터 값을 받아오는 것이 block된 경우
}
```
*/

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Microsecond)

	for {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Microsecond)
		}
	}
}