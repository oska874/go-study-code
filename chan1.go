package main

import "fmt"
import "time"

func main() {

	timeout := make(chan bool, 1)
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // 等待1秒钟
		timeout <- true
	}()
	// 然后我们把timeout这个channel利用起来
	//for {
	select {
	case <-ch:
		// 从ch中读取到数据
		fmt.Print("11\n")
	case <-timeout:
		// 一直没有从ch中读取到数据，但从timeout中读取到了数据
		fmt.Print("22\n")
		close(timeout)
	}
	//}
	_, ok := <-timeout
	if ok == false {
		fmt.Print("closed")
	}
}
