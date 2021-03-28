//先随便选择一个基准值，我们这里选择切片第一个数
//然后循环切片其他值，把比他小的放到左边数组，比他大的放到右边数组，在进行快速排序，
//循环这里n，左右类似二分法logn 所以时间复杂度就是 n*logn
package main

import "fmt"

func kptest(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	value := arr[0]
	var leftArr []int
	var rightArr []int
	for _, i := range arr[1:] {
		if i > value {
			rightArr = append(rightArr, i)
		} else {
			leftArr = append(leftArr, i)
		}
	}
	fmt.Println(leftArr)
	fmt.Println(rightArr)
	pxArr := append(kptest(leftArr), value)
	pxArr = append(pxArr, kptest(rightArr)...)
	return pxArr
}
func main() {
	arr := []int{5}
	a := kptest(arr)
	fmt.Println(a)
}
