package main

import "fmt"

func main() {
	a := []int{0, 0, 1, 2, 2, 3, 4, 5, 4}
	b := xcl(a)
	fmt.Println(b)
}
func reverse(s [5]int) [5]int {
	x := len(s)
	for i := 0; i < x/2; i++ {
		s[i], s[x-i-1] = s[x-i-1], s[i]
	}
	return s
}
func xcl(s []int) []int {
	var x int
	for i := 1; i < len(s); i++ {
		x = s[i-1]
		if s[i] == x {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
