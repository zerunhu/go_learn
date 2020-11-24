package main

import "fmt"

func test(arr []int, size int) int {
	i := 0
	var sum, agv int
	//var agv float32
	for ; i < size; i++ {
		fmt.Println(i)
		sum += arr[i]
	}
	agv = sum / size
	return agv
}

func main() {
	var a [10]int
	var b = [10]int{1, 2, 3}
	var c = []int{1, 3, 4}
	a[0] = 1
	d := b[1]
	fmt.Println(a, b, c, d)

	var e = [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(e)

	g := []int{1, 2, 3}
	f := test(g, 3)
	fmt.Println(f)
}
