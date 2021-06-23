package main

import "fmt"

/*先随便选择数组的一个基准值，我们这里选择切片第一个数，然后循环切片其他值，把比他小的放到左边数组，
比他大的放到右边数组，在分别对左右数组重复上述步骤。
如果每次基准值都选的中位数那么每次左右分组都是对半分，分完整个数组就要logn此循环，每次循环都是便利了数组
所以时间复杂度就是 n*logn，最坏就是每次都选出最大或者最小的，那么就类似于选择排序了 n^2  空间O(N),不稳定
归并排序感觉就是快速排序的最好情况，每次都取中间
快排不稳定是因为根据左右选择的不同总有可能是不稳定的
*/
func kspx(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	index := 0
	var leftArr, rightArr []int
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[index] {
			rightArr = append(rightArr, arr[i])
		} else {
			leftArr = append(leftArr, arr[i])
		}
	}
	pxArr := append(kspx(leftArr), arr[index])
	pxArr = append(pxArr, kspx(rightArr)...)
	return pxArr
}
func main() {
	a := []int{1, 4, 7, 9, 3, 2, 6, 5, 8}
	fmt.Println(kspx(a))
}
