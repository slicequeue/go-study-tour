package main

/*
# 실습: Images
이미 작성했던 picture generator 을 기억하시나요? 이번 시간에는 데이터 슬라이스 대신 image.Image 구현을 반환하는 것을 작성해봅시다.

자신만의 Image 타입을 정의하고, 필요한 메소드 를 구현하고 pic.ShowImage 를 호출합니다.

Bounds 는 image.Rect(0, 0, w, h) 와 같은 image.Rectangle 을 리턴합니다.

ColorModel 은 color.RGBAModel 을 리턴할 것입니다.

At 는 색상을 반환해야 합니다. 마지막 이미지 생성기의 값 v 는 사진의 color.RGBA{v, v, 255, 255} 에 해당합니다.
*/

import "golang.org/x/tour/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
