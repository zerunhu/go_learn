package main

import "fmt"

func main() {
	/* 定义局部变量 */
	a := 10

	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Println("a 小于 20")
	} else {
		fmt.Println("a 大于 20")
	}
	fmt.Printf("a 的值为 : %d\n", a)

	b := 90
	var grade string
	switch b {
	case 80:
		grade = "b"
	case 90:
		grade = "a"
	}
	fmt.Printf("fs is %s\n", grade)

	switch {
	case b < 60:
		fmt.Println("不及格")
	case b > 60:
		fmt.Println("及格")
	}

	//c := 10
	//d := 20
	//select {
	//case c = d:
	//
	//}
	/*
		   select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。

		   select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。
		暂时不清楚
	*/

}
