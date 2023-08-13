package main

/*
Buffered Channels (버퍼가 있는 채널)
Channel은 buffered 될 수 있습니다.
(buffered channel이란 Channel이 버퍼를 가질 수 있다는 의미입니다.)
buffered channel을 초기화하기 위해 make 에 두 번째 인자로 buffer 길이를 제공하십시오.
```
ch := make(chan int, 100)
```
buffered channel 로의 전송은 그 buffer의 사이즈가 꽉 찼을 때에만 block 됩니다. buffer로 부터의 수신은 그 buffer가 비어있을 때 block 됩니다.
buffer가 초과도록 예제를 수정해보고 어떤 일이 발생하는 지 확인해보십시오.
*/

import "fmt"

func main() {
	ch := make(chan int, 2) // 버퍼라는 것은 여기 보이는 두번재 매게변수 2 값 
	// 버퍼라는 것은 2번 받을 수 있다는 것! 
	// 이 설정 없이 무제한적인 받기 허용 보다는 기대하는 수신횟수 보다 많이 받는 경우 제한하여 엄격히 통제 가능함
	ch <- 1
	ch <- 2
	ch <- 3 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}