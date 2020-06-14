package main

import "fmt"

func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)

    customChannel()
}

func customChannel() {
    newChannel:=make(chan string)

    go func()  {
        newChannel<- "another ping"
    }()

    msg:= <-newChannel
    fmt.Println(msg)
}
