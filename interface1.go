package main

import (
	"fmt"
)

type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

/*or*/

type SI interface {
	Get() int
	Put(int)
}

func f(p SI) {
	fmt.Println(p.Get())
	p.Put(1)
}

func f2(p SI) {
	switch t := p.(type) {
	case *S:
	case *R:
	case S: //impossible type switch case: p (type SI) cannot have dynamic type S (missing Get method)
	case R: //impossible type switch case: p (type SI) cannot have dynamic type R (missing Get method)
	default:
	}
}

func main() {

	var ts1 S
	ts1.Put(5)
	fmt.Printf("get %d\n", ts1.Get())

	f(&ts1)

	f2(&ts1)

}
