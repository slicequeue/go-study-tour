package main

/*
# Slice defaults (슬라이스 기본 값)
상한 또는 하한을 생략하면, 슬라이싱할 때 기본 값을 사용할 수 있습니다.
하한의 경우 기본 값은 0이고, 상한의 경우 슬라이스의 길이입니다.
```
var a [10]int
```
위 배열에서 아래 슬라이스 표현식들은 전부 동일한 의미입니다.
```
a[0:10]
a[:10]
a[0:]
a[:]
```
*/

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
