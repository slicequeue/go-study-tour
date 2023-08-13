package main

/*

맵 리터럴
맵 리터럴은 구조체 리터럴과 같지만, 키가 필요합니다.
최상위 타입이 타입 이름일 경우, 리터럴의 요소에서 생략할 수 있습니다.
*/

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		Lat: 40.68433, Long: -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
	"Apple": { // "Vertex" 리터럴 생략 가능
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
