package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	input1 := os.Args
	if len(input1) < 2 {
		fmt.Println("too few parameters")
		return
	}

	fmt.Println(input1[0], "\n", input1[1])
	cmd := exec.Command("ls", "-l")
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("run err ", err)
	//}

	buf, err := cmd.Output()
	if err != nil {
		fmt.Println("out err ", err)
	}

	fmt.Println(buf)

}
