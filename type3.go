package main

import (
	"fmt"
)

//定义接口
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//定义可以用于排序的方法（这样子就可以使用 go 的隐式类型推断）
type Xi []int
type Xs []string

//实现方法
// for number
func (p Xi) Len() int               { return len(p) }
func (p Xi) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xi) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

//for string
func (p Xs) Len() int               { return len(p) }
func (p Xs) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xs) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

//general sort
func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

func main() {
	//number
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	//string
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}

	//sorts
	Sort(ints)
	fmt.Printf("%v\n", ints)
	Sort(strings)
	fmt.Printf("%v\n", strings)

}
