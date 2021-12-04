package main 

import (
	"fmt"
)

func main() {
	// nil interface declaration to string 
	var i interface{} = "hello"

	// assert interface i as string , true , don't panic
	s := i.(string)
	fmt.Println(s)

	// assert interface i as string , ok = true , s = hello
	s, ok := i.(string)
	fmt.Println(s, ok)

	// aserrt interface i as float64 , ok = false, p = 0
	p, ok := i.(float64)
	fmt.Println(p, ok)

	// assert interface i as float64 without ok , so main issue panic
	f := i.(float64)
	fmt.Println(f)
}