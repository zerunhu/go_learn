package main

import "fmt"

type Tree struct {
	value int
	left  int
	right *Tree
}

func main() {
	var a []int
	c, d := testslice(&a, 2)
	fmt.Println(c, d)

	e := retest(1)
	fmt.Println(e)

	f := kebiancs("1", 2, 3, 4)
	fmt.Println(f)
}
func testslice(s *[]int, b int) (*[]int, error) {
	return testerror(s, b)
}
func testerror(s *[]int, b int) (*[]int, error) {
	if b == 1 {
		return nil, fmt.Errorf("error: %d", b)
	}
	*s = append(*s, b)
	return s, nil
}
func retest(a int) (b int) {
	if a == 1 {
		b = a
		return
	} else {
		b = 2
		return
	}
}
func kebiancs(a string, b ...int) int {
	total := 0
	for _, i := range b {
		total += i
	}
	return total
}
