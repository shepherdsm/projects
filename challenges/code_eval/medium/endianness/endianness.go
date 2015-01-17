package main

import (
	"fmt"
	"unsafe"
)

// Endianness
// https://www.codeeval.com/open_challenges/15/
func main() {
	var i int32 = 0x01020304
	u := (*byte)(unsafe.Pointer(&i))

	if *u == 0x04 {
		fmt.Println("LittleEndian")
	} else {
		fmt.Println("BigEndian")
	}
}
