package main

import "fmt"

type name struct {
	val  int
	next *name
}

func main() {
	//aaa := []int{1, 2}
	for i := 5; i > 0; i-- {
		fmt.Println(i)
	}
}
