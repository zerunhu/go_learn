package main

import (
	"fmt"
)

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
进阶：你可以实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案吗？

示例 1：
输入：nums = [1,2,0]
输出：3

示例 2：
输入：nums = [3,4,-1,1]
输出：2

示例 3：
输入：nums = [7,8,9,11,12]
输出：1

*/
/*
第一想法是这样：但是不符合题意：为了减少时间复杂度，可以把 input 数组都装到 map 中，然后 i 循环
从 1 开始，依次比对 map 中是否存在 i，只要不存在 i 就立即返回结果，即所求。 O(n) O(n)

看答案有一个很妙的操作，就是先把所有负数和0先全部改成len(nums)+1,然后操作正整数，如果这个数n在
[1,len(nums)]中间，那么就把数组的index为n-1的数改成-nums[n-1],(这里可能存在重复值并且有可能
会吧原来属于这里的值给改变了，所以有两个地方要注意abs())，这样的话第三遍循环就找数组中第一个不是
负数的index，然后把index+1就是第一个没有出现的正整数了
O(n) O(1)
*/

//func firstMissingPositive(nums []int) int {
//	a := make(map[int]int, len(nums))
//	for _, j := range nums {
//		a[j] = j
//	}
//	for i := 1; i < len(nums)+1; i++ {
//		if _, ok := a[i]; !ok {
//			return i
//		}
//	}
//	return len(nums) + 1
//}
func firstMissingPositive(nums []int) int {
	for i, j := range nums {
		if j <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	for _, j := range nums {
		j = abs(j)
		if j > 0 && j < len(nums)+1 {
			nums[j-1] = -abs(nums[j-1])
		}
	}
	//println(nums[0], nums[1])
	for i, j := range nums {
		if j > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
func main() {
	a := firstMissingPositive([]int{1, 1})
	fmt.Println(a)
}
