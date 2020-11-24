package main

import (
	"fmt"
)

func main() {
	const (
		a = 1
		//如果不赋值后续会自动跟前面一样，iota即索引，<<是左移一位，x<<n == x*(2^n)
		b
		c
		d = 2
		e
		f = 3
		g = iota
		h = 2 << iota
		i
		j
	)
	/*
		变量可变，常量不可变，会报错cannot assign to bbb
		var aaa = 1
		aaa = 2
		const bbb = 1
		bbb = 2
	*/
	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}
