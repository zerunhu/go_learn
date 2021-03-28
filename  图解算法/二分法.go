package main

import "fmt"

func test(arr []int, v int) int {
	button := 0
	top := len(arr) - 1
	var mid int
	for button <= top {
		fmt.Printf("button: %d,top: %d \n", button, top)
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
	l := []int{4, 3, 2, 1}
	msg := test(l, 4)
	fmt.Println(msg)
}
