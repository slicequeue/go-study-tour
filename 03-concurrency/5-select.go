package main

/*
Select
select 문은 goroutine이 다중 커뮤니케이션 연산에서 대기할 수 있게 합니다.
select 는 자의 case들 중 하나가 실행될 때까지 block됩니다.
그리고나서 select문은 해당하는 case를 수행합니다.
만약 다수의 case가 준비되는 경우에는 select가 무작위로 하나를 선택합니다.
*/

import "fmt"


func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, y+x
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
	close(c)
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
