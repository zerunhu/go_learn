package main

import "fmt"

/*
可以将A，B组各自再分成二组。依次类推，当分出来的小组只有一个数据时，可以认为这个小组组内已经达到了有序，
然后再合并相邻的二个小组就可以了。这样通过先递归的分解数列，再合并数列就完成了归并排序。
先递归在合并
递归时间复杂度是logN，合并时间复杂度是N，归并排序时间复杂度NlogN，空间复杂度N，稳定性是稳定的

归并排序是把序列递归地分成短序列，递归出口是短序列只有1个元素（认为直接有序）或者2个序列（1次比较和交换）
然后把各个有序的段序列合并成一个有序的长序列，不断合并直到原序列全部排好序。可以发现，在1个或2个元素时，
1个元素不会交换，2个元素如果大小相等也没有人故意交换，这不会破坏稳定性。那么，在短的有序序列合并的过程
中，稳定是是否受到破坏？没有，合并过程中我们可以保证如果两个当前元素相等时，我们把处在前面的序列的元素
保存在结果序列的前面，这样就保证了稳定性。所以，归并排序也是稳定的排序算法。
*/

func hb(arr1 []int, arr2 []int) []int {
	var arr3 []int
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] > arr2[0] {
			arr3 = append(arr3, arr2[0])
			arr2 = append(arr2[:0], arr2[1:]...)
		} else {
			arr3 = append(arr3, arr1[0])
			arr1 = append(arr1[:0], arr1[1:]...)
		}
	}
	if len(arr1) > 0 {
		arr3 = append(arr3, arr1...)
	} else {
		arr3 = append(arr3, arr2...)
	}
	return arr3
}
func dg(arr []int) []int {
	if len(arr) < 2 { //5,3,2,5   3,5  2,5
		return arr
	}
	arr1 := arr[:len(arr)/2]
	arr2 := arr[len(arr)/2:]
	return hb(dg(arr1), dg(arr2))
}
func main() {
	a := []int{1, 4, 7, 9, 3, 2, 6, 5, 8}
	fmt.Println(dg(a))
}
