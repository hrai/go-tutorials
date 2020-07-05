// Package main provides ...
package main

import (
    "os"
)

func main() {
    panic("a prob")

    _,err:=os.Create("/file")

    if err !=nil{
        panic(err)
    }
}
