package main

import "fmt" 
import "os"

func main() {

    var x int8 = 0
    if x > 0 {
        fmt.Printf("0000\n")
    } else {
        fmt.Printf("~~~~\n")
    }

    if f,err := os.Open("LICENSE") ; err != nil {
        fmt.Printf("222\n")
    } else {
        fmt.Printf("1111\n")
        if f != nil {
            fmt.Printf("3333\n")
        }
    }

    a := [5]int{1,2,3,4,5}

    for i := 0;i<5;i++ {
        fmt.Printf("%d ",a[i])
    }
    fmt.Printf("\n")

    for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
        a[i], a[j] = a[j], a[i]
        fmt.Printf("%d\n",i)
        if  i ==  1 {
            fmt.Printf("9999")
            continue
            fmt.Printf("8888")
        }
    }

    for i := 0;i<5;i++ {
        fmt.Printf("%d ",a[i])
    }
    fmt.Printf("\n")


    i := 0
    for i < 10{
        i ++
        fmt.Printf("aa %d \n",i)
    }
}
