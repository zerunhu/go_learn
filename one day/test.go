/*
变量
基础语法
包的导入
*/
package main

//package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包
import (
	"./Mathdir"
	"fmt"
)

//import "fmt" 告诉 Go 编译器这个程序需要使用 fmt, ./Mathdir(这里是目录) 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
func main() { //需要注意的是 { 不能单独放在一行，所以以下代码在运行时会产生错误
	//func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
	var a, b int = 1, 2
	//定义变量，可以类似于python使用迭代复制 -- 这里的int可以不必声明，会根据后面的值来自动定义类型
	//如果你声明了一个局部变量却没有在相同的代码块中使用它，会得到编译错误--> a declared and not used。但是全局变量是允许声明但不使用的
	//指定变量类型，如果没有初始化，则变量默认为零值
	var c int    //就是0
	var d string //即为"" 空字符串，看输出的结果是一个空的，就是什么也没有
	var e bool   //零值即为false
	/*
		以下几种类型的零 nil：
		var a *int
		var a []int
		var a map[string] int
		var a chan int
		var a func(string) int
		var a error // error 是接口
	*/
	f, g := 1, 2 //这个就等于第一个赋值语句var intVal,intVal1 int = 1, 2
	/*
		还可以通过省略var的方式来定义变量,:= 左侧如果没有声明新的变量，就产生编译错误
		var intVal int
		intVal :=1   这时候会产生编译错误,因为:=是一个声明语句不是一个赋值语句，赋值直接用=
		但是同时，这里的:=赋值只能用于函数中，如果在函数外部使用会报错
	*/
	var ( // 这种因式分解关键字的写法一般用于声明全局变量,不要写在函数中，虽然不会报错
		h int
		i bool
	)
	//_, j = 5, 7
	//_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
	fmt.Println(a, b)
	fmt.Println(c, d, e)
	fmt.Println(f, g)
	fmt.Println(h, i)
	//fmt.Println(...) 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
	//使用 fmt.Print("hello, world\n") 可以得到相同的结果
	fmt.Println(mathtest.Add(1, 2))
	//这里调用了Mathdir目录的mathtest包的Add函数，这里的目录名和包名没有必然联系，这里和python有一点不同
	//python其实import ./a 文件之后，通过a.func()来调用函数，没有包一层的概念

	k := 1
	var l = k
	fmt.Println(k, l, &k, &l)
	/*
		所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：
		当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝：
		可以通过 &i 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。值类型的变量的值存储在栈中。
		内存地址会根据机器的不同而有所不同，甚至相同的程序在不同的机器上执行后也会有不同的内存地址。
	*/
}
