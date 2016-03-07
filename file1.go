package main

import (
"os"
)

func xxx() {
	fd,err := os.Open("./LICENSE")
	defer fd.Close()
	buf := make([]byte, 1024)
	n,_ := fd.Read(buf)
	if n != 0 {
		os.Stdout.Write(buf[:n])
	}
	print(err)
	return
}

func main() {
	xxx()
}

