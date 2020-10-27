package main

import "fmt"

func main() {
	a := 10
	b := "a"
	c := 10
	var d *int
	d = &c

	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&c)
	fmt.Println(*d)
}
