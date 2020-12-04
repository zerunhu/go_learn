package main

import "fmt"

type Value map[string]int
type List []int

func (v Value) Get(key string) int {
	return v[key]
}
func (l List) list() int {
	return l[0]
}
func main() {
	vv := Value{"a": 1, "b": 2}
	a := vv.Get("a")
	fmt.Println(a)
	b := List{1, 2, 3}
	c := b.list()
	fmt.Println(c)

	//var aa string
	type aa string
	aaa := aa("1")
	fmt.Println(aaa)
}
