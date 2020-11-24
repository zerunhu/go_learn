package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "is newline")
var s = flag.String("s", " ", "sep")

func main() {
	fmt.Println(*n, *s)
	flag.Parse() //这里可以吧参数全部获取
	fmt.Println(*n, *s)
	fmt.Println(strings.Join(flag.Args(), *s))
	if !*n {
		//fmt.Println(1)
		fmt.Println()
	}
}
