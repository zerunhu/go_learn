package main

import (
	"time"

	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i < 5; i++ {
			time.Sleep(time.Second * 1)
			fmt.Println(i)
		}
		ch <- 1
	}()
	time.Sleep(time.Second)
	fmt.Println(1)
	<-ch //等获取到值了才会恢复，之前会一直阻塞，就可以让主函数等待
}
