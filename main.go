package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type Demo2 struct {
		m struct{} // 0
		n int8     // 1
		p int
	}

	var d2 Demo2
	fmt.Println(unsafe.Sizeof(d2)) // 2
}
