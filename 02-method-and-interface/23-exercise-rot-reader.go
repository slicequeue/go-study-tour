package main

/*
# 실습: rot13Reader
일반적인 패턴은 다른 io.Reader 을 래핑하여 어떤 방식으로든 스트림을 수정하는 io.Reader 입니다.
예를 들어, gzip.NewReader 함수는 io.Reader (압축된 데이터 스트림) 을 가져와 *gzip.Reader 을 반환합니다. 
또한 io.Reader (압축 해제된 데이터 스트림)도 구현하는 Reader 입니다.

모든 알파벳 문자에 io.Reader 를 구현하고, io.Reader 에서 읽는 rot13Reader 를 구현하고 rot13 에 대체 암호를 적용하여 스트림을 수정합니다.
rot13Reader 타입은 제공됩니다. io.Reader 메소드를 Read 메소드를 활용하여 만듭니다.
*/

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read([]byte) (n int, err error) {
	return 1, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
