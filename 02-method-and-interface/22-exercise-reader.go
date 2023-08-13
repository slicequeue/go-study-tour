package main

/*

실습: Readers
ASCII 문자의 무한 스트림을 방출하는 Reader 유형 구현 'A'.
*/

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}