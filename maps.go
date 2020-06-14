package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, key_exists := m["k2"]
    fmt.Println("key_exists:", key_exists)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

    map1:=make(map[string] int)
    map1["a"]=23
    map1["b"]=223
    fmt.Println(map1)
}
