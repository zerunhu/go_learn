package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				//ch <- i
				fmt.Println("tuichu")
				return
			default:
				time.Sleep(time.Second * 1)
				fmt.Println(i)
				ch <- i
			}
		}
	}()
	for {
		x := <-ch
		//fmt.Println(<-ch)
		//if x == 10 {
		//time.Sleep(time.Second * 5)
		//done <- struct{}{}
		//time.Sleep(time.Second * 5)
		//} else {
		fmt.Println(x)
		//}
		//fmt.Println(<-ch)
	}
}
