package main

import "fmt"

func ReadFull(buf []byte) (n int, err int) {
	n = 9
	err = 3
	return
}

func f() (ret int) {
	defer func(y int) {
		ret = y
	}(6)
	return 0
}

func varfunc(arg1 ...int) {
	for _, y := range arg1 {
		fmt.Printf("%d \n", y)
	}
}

func varfunc2(arg ...int) {
	varfunc(arg...)
	fmt.Printf("\n")
	varfunc(arg[:1]...)
}

func namfunc() {
	a := func() {
		println("Hello")
	}
	a()
}

func ff(n int) {
	fmt.Printf("called %d\n", n)
}
func call_back(y int, f func(int)) {
	f(y)
}
func mapfunc() {
	var xs = map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
		/* ... */
	}

	fmt.Printf("%d %d\n", xs[1](), xs[2]())
}

func f_panic() {
	panic(42)
}

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}

func main() {
	var buf1 []byte

	x, y := ReadFull(buf1)
	fmt.Printf("%d %d\n", x, y)

	x1 := f()

	fmt.Printf("%d \n", x1)

	s1 := make([]int, 3)
	s1 = append(s1, 4)
	varfunc(s1[0], s1[1], s1[3])
	varfunc(s1[0], s1[1], s1[2], s1[3])
	fmt.Printf("\n")
	varfunc2(s1[0], s1[1], s1[2], s1[3])

	namfunc()
	mapfunc()
	call_back(1, ff)

	fmt.Printf("%d\n", throwsPanic(f_panic))
}
