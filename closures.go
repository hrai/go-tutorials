package main

import (
	"fmt"
)

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func returnsClosure() func() int {
    a:=7
    return func() int {
        fmt.Println("calling closure code")
        a=a+2
        return a
    }
}

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())

    increaseNum:=returnsClosure()

    fmt.Println(increaseNum())
    fmt.Println(increaseNum())
    fmt.Println(increaseNum())
    fmt.Println(increaseNum())
    fmt.Println(increaseNum())
}
