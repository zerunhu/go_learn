package main

import (
	"fmt"
)

type WordScan struct {
	world int
	scan  int
}

func (wc *WordScan) Write(p []byte) (n int, err error) {
	wc.world += len(p)
	wc.scan += len(p)
	return len(p), nil
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}
func main() {
	a := "123"
	var c WordScan
	fmt.Fprintf(&c, "a%s", a)
	fmt.Fprintf(&c, "a%s", a)
	fmt.Println(c)
}
