package main

import "fmt"
import "hello"

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
	return int(C.random())
}
func Seed(i int) {
	C.srandom(C.uint(i))
}
func main() {
	Seed(100)
	fmt.Println("Random:", Random())
	hello.Hello()
}
