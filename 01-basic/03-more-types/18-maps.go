package main

/*
# Maps (맵)
맵은 키를 값에 매핑합니다.
맵의 zero value는 nil 입니다. nil 맵은 키도 없고, 키를 추가할 수도 없습니다.
make 함수는 주어진 타입의 초기화되고 사용 준비가 된 맵을 반환합니다.
*/

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}