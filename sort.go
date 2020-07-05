package main

import (
    "fmt"
    "sort"
)

func main() {
    strs:=[]string{"c","b","a"}
    sort.Strings(strs)
    fmt.Println(strs)

    ints:=[]int{32,2,29}
    sort.Ints(ints)
    fmt.Println(ints)

    s:=sort.IntsAreSorted(ints)
    fmt.Println("Sorted:",s)
}

