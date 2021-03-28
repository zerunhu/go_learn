package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个包括n个整数的数组nums和一个目标值target。找出nums中的三个整数，使得它们的和与target最接近。返回这三个数的和。假定每组输入只存在唯一答案。
示例：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。   [-4,-1,1,2]
*/

/*
没啥好说的 和上一题三数之和一样，排序，循环第一个，然后剩下的进行双指针夹逼，吧第一次的做最小值
剩下的比他小就替换最小值，返回就完事了， O(N^2) O(1)
*/
func threeSumClosest(nums []int, target int) int {
	var minCha float64
	var sum int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			cha := (nums[i] + nums[left] + nums[right]) - target
			chaAbs := math.Abs(float64(cha))
			if i == 0 && left == i+1 && right == len(nums)-1 {
				minCha = chaAbs //这里是处理第一次的值，不能直接定义一个最小值
			}
			if chaAbs <= minCha { //等于其实是保证长度为3的数组第一次循环就要返回
				minCha = chaAbs
				sum = nums[i] + nums[left] + nums[right]
			}
			if cha < 0 {
				left++
			} else if cha > 0 {
				right--
			} else {
				return nums[i] + nums[left] + nums[right]
			}
		}
	}
	return sum
}

func main() {
	a := []int{-1, 2, 1, -4}
	b := threeSumClosest(a, 1)
	fmt.Println(b)
}
