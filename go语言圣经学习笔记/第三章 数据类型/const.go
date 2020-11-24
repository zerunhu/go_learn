package main

import "fmt"

func main() {
	const a = 1
	const b = 2
	c := 3
	d := 4
	const e = b + a
	//const f = [3]int{1,2,3}
	fmt.Println(a, b, c, d, e)

	type we int
	const (
		aa we = iota
		bb
		cc
	)
	fmt.Println(aa, bb, cc)

}
