package main

import "fmt"

/*
函数是基本的代码块，用于执行一个任务。
Go 语言最少有个 main() 函数。
你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。
函数声明告诉了编译器函数的名称，返回类型，和参数。
Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

函数定义
Go 语言函数定义格式如下：
func function_name( [parameter list] ) [return_types] {
   函数体
}
函数定义解析：
func：函数由 func 开始声明
function_name：函数名称，函数名和参数列表一起构成了函数签名。
parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
函数体：函数定义的代码集合。
*/
func max(num1 int, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//多个返回
func stx(a string, b string) (string, string) {
	return a, b
}

//闭包
func test(x int) func(int) int {
	f := func(i int) int {
		fmt.Println(i)
		x += i
		return x
	}
	return f
}

//方法
//这里相当于定义了一个类，定义了类的变量，但是类方法是另外定义的
//后续知道这里是叫做结构体的一种东西，相当于定义了一些列键值对把，想要获取结构体的值
//需要类似于实例化的操作，即 var d Dockerapi --> d.host，下面的方法也是如此
type Dockerapi struct {
	host int
	port int
}

//这里是类的方法，d相当于一个临时实例，用来调用类的变量
func (d Dockerapi) getUrl() int {
	return d.host + d.port
}

func main() {
	/* 定义局部变量 */
	var a = 100
	var b = 200
	//var ret int

	/* 调用函数并返回最大值 */
	ret := max(a, b)
	fmt.Printf("最大值是 : %d\n", ret)
	x, _ := stx("1", "2")
	fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(x, y)

	c := test(1)
	fmt.Println(c(1))

	//获取类实例，然后调用类方法
	var d = Dockerapi{1, 21}
	//d.host = 1
	//d.port = 21
	fmt.Println(d.getUrl())
}
