package main

import "fmt"

func main() {
	var a interface{}
	if a.(string) == "" {
		fmt.Printf("a")
	}
	fmt.Println(a)
	fmt.Printf("typea:%T", a)
}
