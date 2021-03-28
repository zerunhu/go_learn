/*
给定两个大小分别为m和n的正序（从小到大）数组nums1,nums2。
请你找出并返回这两个正序数组的中位数。

示例 1:
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例 2:
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

这一题智商不够 只能用O(n+m)的时间复杂度
*/

package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a := append(nums1, nums2...)
	sort.Sort(sort.IntSlice(a))
	n := len(a)
	if n%2 == 0 {
		num := float64(a[n/2]+a[n/2-1]) / 2
		return num
	}
	num := float64(a[(n-1)/2])
	return num
}

func main() {
	num1 := []int{1, 3}
	num2 := []int{2}
	a := findMedianSortedArrays(num1, num2)
	fmt.Println(a)
}
