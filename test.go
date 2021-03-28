package main

import (
	"fmt"

	errors "golang.org/x/xerrors"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
	n int
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
func test(p Phone, x string) {
	fmt.Println(p, x)
}

func main() {
	err := errors.New("abc")

	a := err.Error()
	fmt.Println(err)
	fmt.Println(a)
}
