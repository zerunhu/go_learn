package main

import "fmt"

func main() {
	//defer p1()    在函数返回之前执行，或者panic之前或者函数执行结束，defer倒序返回
	a := test()
	fmt.Println(a)
}
func p1() {
	fmt.Println("1")
}
func p2() {
	fmt.Println("2")
}
func p3() {
	fmt.Println("3")
}
func test() int {
	defer p1()
	defer p2()
	return 111
}
