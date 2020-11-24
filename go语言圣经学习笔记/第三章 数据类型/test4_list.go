package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [4]int{1, 2, 3, 4}
	fmt.Printf("%T,%T", a, b)
	//数组的长度是数组类型的一个部分，因此[3]int和[4]int是两种不同的数组类型,比如我们==两个数组的时候会报错
	c := [...]int{2: 2}
	d := [5]int{2: 2}
	fmt.Println(c, d)
	//可以直接指定索引值来创建数组，没有指定的就是对应数据类型的空值
	f := [3]int{1, 2, 3}
	f[2] = 1
	fmt.Println(f)

	g := [3]int{1, 2, 3}
	h := g
	h[2] = 1
	fmt.Println(g, h)
	//数组不是引用所以修改h的值g不会变，但是slice会变
}
