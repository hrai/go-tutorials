package main

import (
    "fmt"
    "sort"
)

func main() {
    strs:=[]string{"c","b","a"}
    sort.Strings(strs)
    fmt.Println(strs)
}

