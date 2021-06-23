package main

import "fmt"

/*
给定一个整数数组 nums，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。

示例 2：
输入：nums = [1]
输出：1

示例 3：
输入：nums = [0]
输出：0
*/

/*
我们用dp的方法来解，用f(n)代表以n结尾的最大连续数组和，那么可以获得dp的状态转移方程，if f(n-1)>0
f(n)=f(n-1)+nums[n], if f(n-1)<0,f(n)=nums[n]。在一个循环里面求出所有的n结尾的最大连续数组和，
并且在循环中用一个值max代替所有之前的最大连续数组和的最大值并且替换，
O(N) O(1)
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
		if res < nums[i] {
			res = nums[i]
		}
	}
	return res
}
func main() {
	a := []int{1, -2, 1, 3}
	b := maxSubArray(a)
	fmt.Println(b)
}
