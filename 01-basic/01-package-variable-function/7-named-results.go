package main

/*
## 이름이 주어진(기명의) 반환 값
Go의 반환 값은 이름이 정해질 수도 있습니다. 그 경우, 함수의 맨 위에서 정의된 변수처럼 다뤄집니다.
이러한 이름들은 반환 값의 의미를 설명하는 데에 사용되어야합니다.
인자가 없는 return 문은 이름이 주어진 반환 값을 반환합니다. 이것을 "naked" return 이라고 합니다.
여기 이 예제에서처럼 naked return문은 짧은 함수에서만 사용되어야합니다. 긴 함수에서는 그것이 가독성을 떨어뜨릴 수 있습니다.
*/

import "fmt"

func split(sum int) (x, y int, test string) {
	x = sum * 4 / 9
	y = sum - x
	test = "test"
	return // naked return문은 짧은 함수에서만 사용
}

func main() {
	fmt.Println(split(17))
}
