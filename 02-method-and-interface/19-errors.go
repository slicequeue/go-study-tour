package main

/*
# Errors
Go 프로그램은 'error' 값으로 오류 상태를 표현합니다.
error 타입은 fmt.Stringer 와 유사한 내장 인터페이스 입니다:
```
type error interface {
    Error() string
}
```
(fmt.Stringer 와 마찬가지로 fmt 패키지는 값을 출력할 때 error 인터페이스를 찾습니다.)

함수는 종종 error 값을 반환하며, 호출 코드는 오류가 nil 과 같은지 테스트하여 오류를 처리해야합니다.
```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
```
fmt.Println("Converted integer:", i)
nil error 는 성공을 나타냅니다; nil이 아닌 error 는 실패를 나타냅니다.
*/

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}