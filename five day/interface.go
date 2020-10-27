package main

import "fmt"

type phone interface {
	get_num() int
}
type hw struct {
	num int
}
type vivo struct {
	num int
}

func (hw hw) get_num() int {
	return hw.num
}
func (vivo vivo) get_num() int {
	return vivo.num
}
func main() {
	a := hw{3000}
	fmt.Println(a.get_num())
	var b phone
	b = vivo{2000}
	fmt.Println(b.get_num())
}
