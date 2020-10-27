package main

import "fmt"

func Factorial(n int) int {
	if n > 0 {
		result := n * Factorial(n-1)
		return result
	}
	return 1
}
func fblqsl(n int) int {
	if n < 2 {
		return n
	} else {
		return fblqsl(n-1) + fblqsl(n-2)
	}
}
func main() {
	a := Factorial(3)
	fmt.Println(a)
	b := fblqsl(3)
	fmt.Println(b)
}
