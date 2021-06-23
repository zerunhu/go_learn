package main

import "fmt"

func test(arr []int, v int) int {
	button := 0
	top := len(arr) - 1
	var mid int
	for button <= top {
		mid = (top + button) / 2
		if arr[mid] < v {
			button = mid + 1
		} else if arr[mid] > v {
			top = mid - 1
		} else if arr[mid] == v {
			return mid
		}
	}
	return -1
}

func main() {
	l := []int{2, 3, 4}
	msg := test(l, 2)
	fmt.Println(msg)
}
