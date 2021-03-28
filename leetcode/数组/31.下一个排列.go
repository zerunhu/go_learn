package main

import "fmt"

/*
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须 原地 修改，只允许使用额外常数空间。
题目意思一开始没懂，其实就是吧这个nums组成要给数字比如[1,2,3]就是123，然后找到要给大于他
但是理他最近的一个数，就是132，然后返回[1,3,2]

示例 1：
输入：nums = [1,2,3]
输出：[1,3,2]

示例 2：
输入：nums = [3,2,1]
输出：[1,2,3]

示例 3：
输入：nums = [1,1,5]
输出：[1,5,1]

示例 4：
输入：nums = [1]
输出：[1]
*/

/*
这里分为几种情况，长度为0，1的时候不说了 就是不变的，长度大于一的时候：
如果一个数组完全是降序排列就是最大的了就直接最后倒叙返回，
如果不是，那就从后往前找到第一个nums[i]<nums[i+1]的数字，让他让他和后面那些大于他的数字中
最小的数来替换，然后后面的数倒叙即可
[2,5,4,3] 先找到nums[i]<nums[i+1]就是2，然后在和大于他但是最小的数3替换[3,5,4,2]因为后面
的肯定是倒叙的，以你为我们排列过，第一轮判断就是找到第一个nums[i]<nums[i+1]，那么前面的数组
必然是nums[i]>nums[i+1]，即倒叙排列。然后吧倒叙转为正序即可最小

这里我觉得时间是O(n) 这里看着是for循环叠了两层，实际上并没有n^2，因为第一层执行的时候只有
满足某种情况才执行第二层，然后直接推出，实际上吧这两个for循环并列也是一样的
空间： O(1)
*/
func nextPermutation(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	} else {
		fastLeft := len(nums) - 2
		for i := fastLeft; i >= 0; i-- { //第一轮循环找到小于
			if nums[i] < nums[i+1] {
				for j := len(nums) - 1; j > i; j-- { //第二轮循环找到大于但是最小的
					if nums[j] > nums[i] {
						nums[j], nums[i] = nums[i], nums[j]
						break
					}
				}
				swap(nums[i+1:])
				return nums
			}
		}
		swap(nums)
		return nums
	}
}
func swap(arr []int) []int {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}
func main() {
	nums := []int{5, 1, 1}
	fmt.Print(nextPermutation(nums))
}
