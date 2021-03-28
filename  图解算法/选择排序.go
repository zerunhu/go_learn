package main

//找出最大的 然后插入新数组
import "fmt"

func findBig(arr []int) (int, int) {
	var num int
	var index int
	for i, j := range arr {
		if j > num {
			num = j
			index = i
		}
	}
	return index, num
}
func main() {
	arr := []int{4, 3, 6, 2, 7, 1}
	sortArr := make([]int, 0, len(arr))
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		index, bigNum := findBig(arr)
		arr = append(arr[:index], arr[index+1:]...)
		sortArr = append(sortArr, bigNum)
	}
	fmt.Println(sortArr)
}
