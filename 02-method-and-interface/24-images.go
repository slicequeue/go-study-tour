package main

/*
# Images
Package image 는 Image 인터페이스를 정의합니다:
```
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```
참고: Bounds 메소드의 Rectangle 반환 값은 선언이 내부에 있으므로 실제로는 image 패키지에 속한 선언인 image.Rectangle 입니다.
(자세한 내용은 the documentation 을 참고하세요. )
color.Color 및 color.Model 타입도 인터페이스이지만, 사전 정의된 구현 `color.RGBA`및 color.RGBAModel 을 사용하여 이를 무시합니다. 
이러한 인터페이스 및 타입은 image/color package 에 자세히 서술되어있습니다.
*/

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
