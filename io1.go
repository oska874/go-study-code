package main

import "os"
import "fmt"
import "bufio"

func noio1() {
	buf := make([]byte, 1024)
	f, _ := os.Open("./io1.go")

	defer func() {
		fmt.Println("bye & close")
		f.Close()
	}()

	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

func io1() {
	buf := make([]byte, 1024)
	f, _ := os.Open("")

	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[0:n])
	}
}

func main() {
	noio1()
	io1()
}
