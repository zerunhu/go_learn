package main

import "fmt"

type Point struct {
	x, y int
}
type Path []Point

func main() {
	a := Point{1, 2}
	fmt.Println(a.test(a))
	b := Point{1, 2}
	var pa Path
	c := append(pa, a, b)
	fmt.Println(c)
	var d []Point
	d = append(d, a)
	fmt.Println(d)
}
func (p Point) test(q Point) int {
	return p.x + p.y + q.x
}
