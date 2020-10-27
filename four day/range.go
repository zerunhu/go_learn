package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	num := 0
	for _, num := range nums {
		fmt.Println(num)
		num += num
		fmt.Println(num)
	}
	fmt.Println(num)

	anums := []int{2, 3, 4}
	asum := 0
	for _, anum := range anums {
		asum += anum
	}
	fmt.Println("sum:", asum)
}
