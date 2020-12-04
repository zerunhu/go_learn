package main

import (
	"time"

	"fmt"
)

func test3(a chan time.Duration, i time.Duration) <-chan time.Duration {
	go func() {
		time.Sleep(time.Second * i)
		fmt.Println(i, "ready")
		a <- i
	}()
	if i == 3 {
		time.Sleep(time.Second)
	} else {
		time.Sleep(time.Second * 2 * i)
	}
	//time.Sleep(time.Second * 2 * i)
	fmt.Println(i)
	return a
}
func main() {
	a := make(chan time.Duration, 10)
	b := make(chan time.Duration, 10)
	c := make(chan time.Duration, 10)
	//time.Sleep(time.Second * 10)
	select {
	case <-test3(a, 1):
		fmt.Println(111)
		//fmt.Println(d)
	case <-test3(b, 2):
		fmt.Println(222)
	case <-test3(c, 3):
		fmt.Println(333)
	default:
		fmt.Println("default")
	}
}
