package main

import (
	"C"
)
import "fmt"

//export HelloWorld
func HelloWorld(input *C.char) *C.char {
	str := C.GoString(input)
	fmt.Println(str)
	return C.CString("Hello world")
}

func main() {}
