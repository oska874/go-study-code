package main

import (
	"fmt"
	"runtime"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is busy")
	c <- 1
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(4))
	c = make(chan int)
	go ready("t1", 2)
	go ready("t2", 1)
	fmt.Println("im waiting")
	//time.Sleep(5 * time.Second)
	<-c //recve data from channel
	<-c
}
