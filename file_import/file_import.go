package main

import (
	"fmt"
)

func main() {
	TestLog()
}

//TestLog access point
func TestLog() {
	in := [...]float64{10, 6, 3}
	// l := in[0]
	// a := in[1]
	// b := in[2]
	// log.Printf("%#v", in)
	// log.Printf("%#v", l)
	fmt.Printf("%#v", in)
}
