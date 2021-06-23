package main

import "fmt"

/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums在预先未知的某个下标 k（0 <= k < nums.length）上进行了旋转，使数组变为
[nums[k],nums[k+1],...,nums[n-1],nums[0],nums[1],...,nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2] 。
给你旋转后的数组nums和一个整数target，如果nums中存在这个目标值target,则返回它的索引，否则返回-1。

示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1
*/

/*
题目意思就是在数组中找到某个值，其实O（n）就解决了，但是他要求O（logN）的时间复杂度，一般就是
二分法了，正好他是一个有序数组，和普通有序数组二分法唯一的差别就是他是有序数组根据某个值进行了
选择排序重新排列了一下，就是这是一个有两个有序子数组的数组，
[4,5,6,7,0,1,2]
*/

func search(nums []int, target int) int {
	right := len(nums) - 1
	left := 0
	for right >= left {
		middle := (right + left) / 2
		if nums[middle] > nums[right] {
			//这里指前面全是有序数组
			if nums[left] <= target && target <= nums[middle] {
				index := erfenfa(nums[:middle+1], target)
				return index
			} else {
				left = middle + 1
			}
		} else if nums[middle] < nums[right] {
			//这里指middle右边全是小的那个有序数组  5, 6, 0, 1, 2, 3, 4  [0,1,2]
			if target <= nums[right] && target >= nums[middle] {
				index := erfenfa(nums[middle:right+1], target)
				if index == -1 {
					return index
				} else {
					return index + middle
				}
			} else {
				right = middle - 1
			}
		} else {
			if target == nums[middle] {
				return middle
			} else {
				return -1
			}
		}
	}
	return -1
}
func erfenfa(arr []int, target int) int {
	fmt.Println(arr, target)
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (left + right) / 2
		if arr[middle] > target {
			right = middle - 1
		} else if arr[middle] < target {
			left = middle + 1
		} else if arr[middle] == target {
			return middle
		}
	}
	return -1
}
func main() {
	a := []int{8, 9, 2, 3, 4}
	b := search(a, 9)
	fmt.Println(b)
}
