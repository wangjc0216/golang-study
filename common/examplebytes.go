package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main33(){
	hallo:="test1"
	fmt.Printf("%p\n",[]byte(hallo))
	world := hallo
	fmt.Printf("%p\n",[]byte(world))
	hallo = "test2"
	fmt.Printf("%p\n",[]byte(hallo))
	fmt.Printf("%p\n",[]byte(world))
}
func stringtoslicebytetmp(s string) []byte {
	str := (*reflect.StringHeader)(unsafe.Pointer(&s))
	ret := reflect.SliceHeader{Data: str.Data, Len: str.Len, Cap: str.Len}
	return *(*[]byte)(unsafe.Pointer(&ret))
}

func main()  {
	str := "hello"
	by := stringtoslicebytetmp(str)

	by[0] = 'H'
}