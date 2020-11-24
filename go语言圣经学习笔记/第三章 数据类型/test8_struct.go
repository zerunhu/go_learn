package main

import "fmt"

func main() {
	type Emp struct {
		id, code int
		name     string
	}
	var a Emp
	a.id = 1
	b := Emp{1, 2, "a"}
	c := Emp{id: 1, code: 2, name: "a"}
	fmt.Println(a, b, c)

	var aaa *Emp
	aaa = new(Emp) //需要用new函数分配新空间
	aaa.id = 1
	fmt.Println(aaa)

	d := sale(P{1, 2}, 3)
	fmt.Println(d)

	var e Q
	e.P.a = 1 //匿名成员的名称默认为结构类型的名称，这样就导致一个结构体不能包含两个类型一样的结构体
	// 但是匿名类型不能使用简短语句定义
	e.b = 1
	e.c = 1
	fmt.Println(e)
}

type P struct {
	a, b int
}
type Q struct {
	P //结构体嵌套，但是会有很多x.x.x，这里可以使用匿名成员，无需指定名字，直接使用类型，就可以在调用的时候直接指定最底层的变量
	c int
}

func sale(p P, x int) P {
	return P{p.a * x, p.b * x}
}
