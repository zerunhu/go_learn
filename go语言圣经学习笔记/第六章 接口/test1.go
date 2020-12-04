package main

import "fmt"

type Chicken struct {
}

func (c Chicken) IsChicken() bool {
	fmt.Println("我是小鸡")
	return true
}
