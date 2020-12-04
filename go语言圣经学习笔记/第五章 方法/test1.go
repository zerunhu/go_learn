package main

import "fmt"

type Pointt struct {
	x, y int
}

func (p Pointt) testt(n int) Pointt {
	p.x *= n
	p.y *= n
	return p
}
func main() {
	a := Pointt{1, 2}
	//b := &a
	c := a.testt(2)
	fmt.Println(a, c)
}
