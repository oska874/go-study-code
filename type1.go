package main

import "fmt"

func main() {

	var str string = "hello world"
	fmt.Printf(str + "\n")

	s := []rune(str)
	s[0] = '1'
	s2 := string(s)
	fmt.Printf(s2 + "\n")

	var c complex64 = 5 + 5i
	fmt.Printf("Value is: %v\n", c)

	var a int32 = 0x12345678
	var a2 int64 = 0x12345678abcd1234
	var a3 uint16 = 0xf1f2

	fmt.Printf("%x %x %x\n", a, a2, a3)

	var a4 [3]int32
	a4[0] = 0
	a4[1] = 1
	a4[2] = 2
	print(a4[2], "\n")

	a5 := [...]int8{1, 2, 3, 4, 5}
	print(a5[2], "\n")

	sl := make([]int, 10)

	sl[9] = 9
	sl = append(sl, 10)

	fmt.Printf("%d\n", sl[10])

	mon := map[string]int{"mon": 1}
	fmt.Printf("%d \n", mon["mon"])
	mon2 := make(map[string]int)
	mon2["abc"] = 23
	fmt.Printf("%d \n", mon2["abc"])
	mon2["abc12"] = 2333
	fmt.Printf("%d \n", mon2["abc12"])
	delete(mon2, "abc")
	fmt.Printf("%d \n", mon2["abc"])
	v, ok := mon2["abc"]
	fmt.Printf("%d ", v, ok)

	strs := [...]string{"asSASA", "ddd", "dsjkdsjs", "asSASA", "ddd", "dsjkdsjs", "dk"}
	for k, str1 := range strs {
		fmt.Printf("%d %s len %d\n", k, str1, len(str1))
	}

	var sl1 = make([]int, 10)
	fmt.Print(sl1[0], sl1[9])
}
