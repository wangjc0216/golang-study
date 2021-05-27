package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	cpuprofile := "test_profile"
	f, err := os.Create(cpuprofile)
	defer f.Close()
	if err != nil {
		fmt.Println("os create occur error:", err)
		return
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	a := 0

	for i := 0; i < 10000; i++ {
		a += i
		time.Sleep(time.Millisecond)
	}
	fmt.Println(a)
}
