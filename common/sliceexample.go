package main

import (
	"fmt"
)

func main31(){
	a := []int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(a)
	fmt.Println(a[11:])
	a = append(a[:9],a[10:]...)
	fmt.Println(a)

}

