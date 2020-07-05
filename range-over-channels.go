package main

import "fmt"

func main() {

    queue:=make(chan string, 4)
    queue<-"one"
    queue<-"two"
    queue<-"last"
    close(queue)

    for elem:=range queue {
        fmt.Println(elem)
    }
}
