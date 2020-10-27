package main

import "fmt"

func main() {
	var a []int
	var b = []int{1, 2, 3}
	var c = make([]int, 3, 5)
	c = append(c, 4)
	a = append(a, 4)
	b = append(b, 4)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	slice1 := []int{1, 2, 3, 4, 5}
	//var slice2 []int
	//copy要定义一样的空间，如果只是初始化 并不会复制，还是一个[]
	slice2 := make([]int, 5)
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
}
