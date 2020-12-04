package main

import "fmt"

type phone interface {
	call()
}
type iphone struct {
}

func (iphone iphone) call() {
	fmt.Println("iphone")
}

type hw struct {
}

func (hw hw) call() {
	fmt.Println("hw")
}

func main() {
	var p phone
	p = hw{}
	p.call()
}
