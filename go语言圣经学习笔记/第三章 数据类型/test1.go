package main

import (
	"fmt"
	"strings"
)

func main() {
	a := 0
	b := 0
	if a == 0 || b == 1 {
		fmt.Println("a")
	}

	//c := "hello,world"
	//for i, j := range c {
	//	fmt.Printf("%d\t%q\t%d\n", i, j, j)
	//}
	c := "ho/wo/he"

	for i := len(c) - 1; i >= 0; i-- {
		fmt.Printf("第%d次循环\t", i)
		fmt.Println(c)
		if c[i] == '/' {
			c = c[i+1:]
			fmt.Println(c)
			break
		}
	}

	d := "helloworld"
	fmt.Println(strings.LastIndex(d, "w"))

	e := "abc"
	f := []byte(e)
	g := string(b)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%d", g)

}
