package main

import (
	"fmt"
	"os"
	"time"
)

func test7(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	return
}
func main() {
	ch := make(chan int)
	go test7(ch)
	time.Sleep(time.Second * 2)
	for {
		select {
		case msg := <-ch:
			fmt.Fprintln(os.Stdout, msg)
		default:
			fmt.Println("jieshu")
			return
		}
	}

}
