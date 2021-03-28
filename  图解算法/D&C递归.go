package main

import "fmt"

//编写一个递归函数来计算列表的总和sum
func dctest1(arr []int) int {
	popNum := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	if len(arr) == 0 {
		return popNum
	} else {
		return popNum + dctest1(arr)
	}
}

//编写递归函数找出数组最大的数
func dctest2(arr []int) int {
	popNum := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	if len(arr) == 0 {
		return popNum
	}
	v := dctest2(arr)
	if popNum > v {
		return popNum
	} else {
		return v
	}
}

//递归函数实现二分法
func dctest3(arr []int, value int) int {
	if len(arr) == 1 {
		if arr[0] == value {
			return 123
		} else {
			return -1
		}
	}
	middile := (len(arr) - 1) / 2
	if arr[middile] > value {
		arr = arr[:middile-1]
		return dctest3(arr, value)
	} else if arr[middile] < value {
		arr = arr[middile+1:]
		return dctest3(arr, value)
	} else {
		return 123
	}
}
func main() {
	arr := []int{1, 2, 3, 4}
	a := dctest3(arr, 3)
	fmt.Println(a)
}
