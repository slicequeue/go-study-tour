package main

/*
# 실습: Stringers
IPAddr 타입이 fmt.Stringer 을 구현하여 dotted quad로 출력하도록 합니다.
예를 들어, IPAddr{1, 2, 3, 4} 는 "1.2.3.4" 와 같이 출력되어야합니다.
*/

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr. - JOBS DONE!
func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

/*
궁금증! 
func (i *IPAddr) String() string {
- 이렇게 하면 `*IPAddr` 인데 차이가 있을까? 
- 결과는 동일함! 
*/

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
