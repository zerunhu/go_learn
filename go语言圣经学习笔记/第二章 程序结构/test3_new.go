package main

import "fmt"

/*
五种定义
var a int      定义一个变量 不赋值，默认是0值
var a = 1      定义一个值 不需要显式表明类型，会根据值自己判断
a := 1         只能在函数内部定义
a := new(int)  定义的是一个指针, 如果不是这样定义的话，var i *int, *int = 1 将会报错
a = 1          赋值语句不是定义语句
*/
var global *int
var global1 int

func fun1() {
	var x int
	x = 1
	global = &x
}
func p() {
	//var x int
	//x = 1
	global1 += 1 //这里和python不一样，python在函数里是读取不到全局变量的。
}
func ad() {

}
func g() {
	y := new(int)
	*y = 1
}

func main() {
	//p := new(int)
	//fmt.Println(*p)

	fun1()
	p()
	g()
	fmt.Println(*global)
	fmt.Println(global1)
	//fmt.Println(*p)
	//fmt.Println(*y)
}
