package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}
	c <- sum
	//如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。
	fmt.Println("c接受了") //这里输出一个值可以来测试通道的阻塞问题
}

func sum2(s []int, c chan int) {
	for _, i := range s {
		c <- i
	}
	close(c)
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := make(chan int, 3)
	//如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
	go sum(a, c)
	go sum(b, c)
	x := <-c
	time.Sleep(time.Second)
	y := <-c
	fmt.Println(x, y)

	d := make(chan int, 3)
	go sum2(a, d)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 3 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 3 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 4 个数据的时候就阻塞了。实际执行情况是会报错
	for i := range d {
		fmt.Println(i)
	}
}
