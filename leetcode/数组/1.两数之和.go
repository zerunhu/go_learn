/*
题目：
给定一个整数数组 nums和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

package main

import "fmt"

/*
暴力破解
时间复杂度：O(N^2)，其中 NN 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
空间复杂度：O(1)
*/
func twoSum(nums []int, target int) []int {
	for i, j := range nums {
		for x, y := range nums[i+1:] {
			if j+y == target {
				return []int{i, i + x + 1}
			}
		}
	}
	return nil
}

/*
哈希表
时间复杂度：O(N)，其中 NN 是数组中的元素数量。对于每一个元素 x，我们可以 O(1)O(1) 地寻找 target - x。
空间复杂度：O(N)，其中 NN 是数组中的元素数量。主要为哈希表的开销。
*/
func twoSumH(nums []int, target int) []int {
	ha := make(map[int]int)
	for i, j := range nums {
		if value, ok := ha[target-j]; ok {
			return []int{value, i}
		}
		ha[j] = i
	}
	return nil
}
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := twoSumH(a, 8)
	fmt.Println(b)
}
