package main

/*
# Readers
io 패키지는 데이터 스트림의 읽기를 나타내는 io.Reader 인터페이스를 지정합니다.
Go 표준 라이브러리에는 파일, 네트워크 연결, 압축기, 암호 등을 포함하여 인터페이스의 많은 구현 이 포함되어 있습니다.
io.Reader 인터페이스에는 Read 메소드가 있습니다:
```
func (T) Read(b []byte) (n int, err error)
```
Read 는 주어진 바이트 조각을 데이터로 체우고 채워진 바이트 수와 오류 값을 반환합니다. 스트림이 종료되면 io.EOF 오류를 반환합니다.
예제 코드는 strings.Reader 에서 생성할 수 있으며 이는 한번에 8바이트 씩 사용합니다.
*/

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

