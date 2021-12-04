package main

import (
	"fmt"
	"strings"
	"strconv"
)
// new type IpAddr is cap4 array of byte
type IpAddr [4]byte

func (ip IpAddr) String() string {
    s := make([]string, 0, len(ip))
    for _, i := range ip {
        s = append(s, strconv.Itoa(int(i)))
    }
    return strings.Join(s, ".")
}

func main() {
	// define hosts is map key=string, value is IpAddr Interface
	hosts := map[string]IpAddr{
		"loopback": {127,0,0,1},
		"googleDNS": {8,8,8,8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n",name, ip)
	}
}