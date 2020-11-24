package main

import "fmt"

func main() {
	a := 1
	b := 1
	if a == 1 {
		b := 2
		//和python不一样的是这里在if里面定义的变量是在外面无法获取的,如果在内部用=则是取外部的变量,因为python没有定义和赋值的区别导致的
		fmt.Println(b)
	}
	fmt.Println(b)

	c := "hellp"
	for _, c := range c {
		fmt.Printf("%c", c)
		c := c
		fmt.Printf("%c", c)
	}
	fmt.Println(c)

	d := 1
	e := [3]int{1, 2, 3}
	for _, i := range e {
		d = i
	}
	fmt.Println(d)
}
