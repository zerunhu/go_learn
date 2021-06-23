package main

import (
	"time"
)

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

type Test struct {
	a int
}

func (t Test) add() int {
	return t.a * 2
}

type Test2 struct {
	Test
}

func test111(arr []int) {
	arr = append(arr[:3], arr[4:]...)
}
func test11(arr []int) {
	//a := []int{6, 7}
	arr = append(arr, 1)
}

type al struct {
	a int
	b string
}

var TokenExpireDuration = time.Hour * 1

func main() {
}
