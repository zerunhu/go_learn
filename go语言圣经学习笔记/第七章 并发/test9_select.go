package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("start")
	tick := time.Tick(time.Second)
	for c := 10; c > 0; c-- {
		fmt.Println(c)
		<-tick
	}
	about := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		about <- struct{}{}
	}()
	fmt.Println("fas")
	//for {
	//	fmt.Println(<-tick)
	//}
}
