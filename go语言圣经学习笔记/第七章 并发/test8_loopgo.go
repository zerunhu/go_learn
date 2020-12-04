package main

import (
	"time"

	"fmt"
)

//如何理解阻塞，如果没有接受chan的话，进入并发后，他还没来得及执行完就会因为主程序执行完成退出，然后
//并发程序就会一并退出，如果有的话，接受chan的主进程会一直等待，直到并发程序执行完，传去chan值就会继续执行主进程
func test1() {
	time.Sleep(time.Second * 3)
}
func test2(ch chan int) int {

	a := [5]int{1, 2, 3, 4, 5}
	for _, i := range a {
		go func(i int) {
			fmt.Println("x", i)
			test1()
			fmt.Println("wobeizusele", i)
			ch <- i
			fmt.Println("wo  tuichule", i)
		}(i)
	}
	for range a {
		x := <-ch
		fmt.Println("x:", x)
		if x == 2 {
			fmt.Println("ch值", x)
			fmt.Println("ch长度：", len(ch))
			return x
		}
	}
	return 0
}
func main() {
	ch := make(chan int)
	test2(ch)
	time.Sleep(time.Second * 5)
	fmt.Println(ch)
	fmt.Println(len(ch))
	time.Sleep(time.Second * 5)
}
