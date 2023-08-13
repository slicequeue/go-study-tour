package main

/*
# 슬라이스 연습하기
Pic 을 구현해봅시다. 이것은 dy 슬라이스의 길이와 dx 슬라이스의 각 요소를 8 비트 부호 없는 정수로 반환해야 합니다. 
이 프로그램을 실행하면, 이것은 grayscale (well, bluescale) 값으로 변환된 그림을 보여줍니다.
이미지의 선택은 여러분에게 달려 있습니다. 흥미로운 함수로 (x+y)/2 , x*y , 그리고 x^y 를 제공합니다.
(루프를 사용하여 [][]uint8 안에 각 []uint8 을 할당해야 합니다.)
(타입 간에 변환하려면 uint8(intValue) 을 사용하세요.)
*/

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	return [][]uint8 {{1,2}}
}

func main() {
	pic.Show(Pic)
}