package main

import "fmt"

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。
找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。
你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
*/

/*
这一题是经典的二分搜索变种题。二分搜索有 4 大基础变种题：
查找第一个值等于给定值的元素
查找最后一个值等于给定值的元素
查找第一个大于等于给定值的元素
查找最后一个小于等于给定值的元素
https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0034.Find-First-and-Last-Position-of-Element-in-Sorted-Array/
*/
func searchLeftValue(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
func searchRightValue(nums []int, target int) int { //22
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
func searchRange(nums []int, target int) []int {
	left := searchLeftValue(nums, target)
	right := searchRightValue(nums, target)
	return []int{left, right}
}

func main() {
	a := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(a, 9))
}
