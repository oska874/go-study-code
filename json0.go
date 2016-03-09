package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string
	//Authors []string
	Publisher   string
	IsPublished bool
	Price       int
}

func main() {

	gobook := Book{
		"Go语言编程",
		//["XuShiwei", "HughLv","Pandaman", "GuaguaSong", "HanTuo", "BertYuan","XuDaoli"],
		"ituring.com.cn",
		true,
		9}

	gonebook := Book{
		"abc", "def", false, 8}

	b, err1 := json.Marshal(gobook)
	if err1 == nil {
		fmt.Printf("%v\n", string(b))

	}

	c, err2 := json.Marshal(gonebook)
	if err2 == nil {
		fmt.Printf("%v\n", string(c))

	}

}
