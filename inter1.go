package main 

import "fmt"

type inte int

func (a inte) Less(b inte) bool{
    return a<b
}

func (a *inte) Add(b inte) {
    *a+=b
}

type LessAdder interface {
    Less(b inte) bool
    Add(b inte)
}

func main() {
    var a inte = 1
    var b LessAdder = &a

    fmt.Println(b)
}