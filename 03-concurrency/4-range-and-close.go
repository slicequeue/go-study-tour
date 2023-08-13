package main2

/*
Range와 Close
전송자는 더 이상 보낼 데이터가 없다는 것을 암시하기 위해 channel을 close할 수 있습니다. 
수신자는 수신에 대한 표현에 두 번째 매개변수를 할당함으로써 채널이 닫혔는지 테스트할 수 있습니다.
```
v, ok := <-ch
```
만약 더 수신할 값이 없고, channel이 닫혀있다면 ok 는 false 입니다.
for i := range c 반복문은 channel이 닫힐 때까지 반복해서 channel에서 값을 수신합니다.
주의: 절대로 수신자가 아닌 전송자만이 channel을 닫아야합니다. 닫힌 channel에 전송하는 것은 panic을 야기할 것입니다.
하나 더 주의: Channel은 파일과 다릅니다. 당신은 file과 달리 보통 channel을 닫을 필요는 없습니다. 
channel을 닫는 것은 range 반복문을 종료시키는 것과 같이 수신자가 더 이상 들어오는 값이 없다는 것을 알아야하는 경우에만 필요합니다.
*/

import "fmt"


func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}