package main

import (
	"fmt"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v ", i[0], i[1], i[2], i[3])
}

func main() {
	addrs := map[string]IPAddr{
		"lookback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v : %v \n", n, a)
	}
}
