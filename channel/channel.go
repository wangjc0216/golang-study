package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var exit = make(chan string, 1)

func main() {
	go dealSignal()
	exited := make(chan struct{}, 1)
	go channel1(exited)
	count := 0
	t := time.Tick(time.Second)
Loop:
	for {
		select {
		case <-t:
			count++
			fmt.Printf("main run %d\n", count)
		case <-exited:
			fmt.Println("main exit begin")
			break Loop
		}
	}
	fmt.Println("main exit end")
}

//func dealSignal() {
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
//	go func() {
//		<-c
//		exit <- "shutdown"
//	}()
//}

func dealSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	exit <- "shutdown"
}

func channel1(exited chan<- struct{}) {
	t := time.Tick(time.Second)
	count := 0
	for {
		select {
		case <-t:
			count++
			fmt.Printf("channel1 run %d\n", count)
		case <-exit:
			fmt.Println("channel1 exit")
			close(exited)
			return
		}
	}
}
