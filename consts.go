package main

import (
    "fmt"
    "math"
)

const s string = "constant"
const pi = 3.14

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(pi/10)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))
    fmt.Println(math.Ceil(n))
}
