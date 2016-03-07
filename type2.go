package main

import "fmt"

type NameAge struct {
	name string
	age  int
}

func main() {
	a := new(NameAge)
	a.name = "Pete"
	a.age = 42
	fmt.Printf("%v\n", a)
	t22()

	typ2()
}

func doSomething1(n1 *NameAge, n2 int) { fmt.Printf("%d\n", n2) }

func (n1 *NameAge) doSomething2(n2 int) { fmt.Printf("%d\n", n2*n2) }

func t22() {
	var n *NameAge
	n.doSomething2(4)

	doSomething1(n, 4)
}

type Base struct {
	Name string
	Age  int
}

type Child struct {
	Base // 匿名字段， 默认把Base的所有字段都继承过来了。 这样看起来才像真正的继承
	Age  int
}

type foo struct{ int }

func typ2() {
	c := new(Child)
	c.Name = "hello" // 可以直接使用Base中的字段
	c.Age = 20       // 如果有重复的， 则最外的优先
	c.Base.Age = 21

	fmt.Println(c.Name)     // hello
	fmt.Println(c.Age)      // 20
	fmt.Println(c.Base.Age) // 要访问Base中的，可以这样写 0

	c1 := new(foo)
	c1.int = 12
	fmt.Print(c1)
}
