package main

import "fmt"

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，
返回它将会被按顺序插入的位置。你可以假设数组中无重复元素。

示例 1:
输入: [1,3,5,6], 5
输出: 2

示例2:
输入: [1,3,5,6], 2
输出: 1

示例 3:
输入: [1,3,5,6], 7
输出: 4

*/

/*
没啥好说的，就是二分法的变形：
查找第一个大于等于给定值的元素
查找最后一个小于等于给定值的元素
就查夹在mid和mid+1或者mid-1中间的数，和边界情况mid在数组第一个或者最后一个
*/
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		} else if nums[mid] < target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid + 1
			}
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
func main() {
	a := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(a, 5))
}
