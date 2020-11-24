package main

import "fmt"

func f(p *int) int {
	*p++
	return *p
}

func ff(p int) int {
	p++
	return p
}
func main() {
	v := 1
	p := f(&v)
	q := f(&v)
	x := 1
	xx := ff(x)
	fmt.Println(v, p, q)
	fmt.Println(x, xx)
}
