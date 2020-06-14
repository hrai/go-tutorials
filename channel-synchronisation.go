package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done

    channel := make(chan string, 1)
    go newWorker(channel)
    fmt.Println(<-channel)
}

func newWorker(channel chan string) {
    fmt.Println("work started...")
    time.Sleep(time.Second*2)
    fmt.Println("work done")

    channel <- "going out of newWorker"
}
