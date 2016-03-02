package main

import "fmt"

func main() {

	var str string = "hello world"
	fmt.Printf(str+"\n")

	s := []rune(str)
	s[0] = '1'
	s2 := string(s)
	fmt.Printf(s2+"\n")

	var c complex64 = 5+5i
	fmt.Printf("Value is: %v\n", c)

	var a int32 = 0x12345678
	var a2 int64 = 0x12345678abcd1234
	var a3 uint16 = 0xf1f2

	fmt.Printf("%x %x %x\n",a,a2,a3)
}
