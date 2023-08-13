package main

/*
# 
Goroutines
goroutine 은 Go 런타임에 의해 관리되는 경량 쓰레드입니다.
```
go f(x, y, z)
```
새로운 goroutine을시작합니다.
```
f(x, y, z)
```
f와 x, y, z의 evaluation은 현재의 goroutine에서 일어나고, f 의 실행은 새로운 goroutine에서 일어납니다.

goroutine은 같은 주소의 공간에서 실행되고, 따라서 공유된 메모리는 synchronous(동기적) 해야합니다. 
Go에 다른 기본형들이 존재하므로 당신이 Go에서 sync와 관련된 기본 기능이 필요없다하더라도 sync 패키지는 유용한 기본형을 제공합니다. 
(다음 페이지에서 자세히 확인해보십시오.)
*/

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i, s)
	}
}

func main() {
	go say("world")
	say("hello")
}

