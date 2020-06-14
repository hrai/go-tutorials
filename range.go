package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }

    test()
}

func test() {

    nums:= []int{23,233}
    for i, n:= range nums {
        fmt.Println(i,n)
    }

    dict:=map[string]string{"a":"area","b":"build"}
    for k,v:= range dict{
        fmt.Println("key is", k)
        fmt.Println("value is", v)
    }
}
