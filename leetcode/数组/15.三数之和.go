/*
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例 2：
输入：nums = []
输出：[]
*/

package main

import (
	"fmt"
	"sort"
)

//没啥说的暴力破解来一遍  O(N^3)
func threeSum(nums []int) [][]int {
	var arr [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if i+j+k == 0 {
					arr = append(arr, []int{i, j, k})
				}
			}
		}
	}
	return arr
}

//就先排序，然后第一轮正常和暴力一样，第二轮开始双指针，这里双指针和普通的有点不一样，左一直动，
//右边如果动防止左边动就退一格，主要是因为正常的左右互移不太好吧同样的值去重，所以就这样了，类似
//O(N^2)排序是O(nlogn),所以最后是O(N^2)，O(n)
func threeSum1(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for j := left; j < right; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := nums[i] + nums[j] + nums[right]
			if sum > 0 {
				right--
				j--
			} else if sum < 0 {
				continue
			} else {
				result = append(result, []int{nums[i], nums[j], nums[right]})
			}
		}
	}
	return result
}

func main() {
	a := []int{-1, 0, 1, 2, -1, -4}
	//a := []int{0, 0, 0, 0}
	//[-4,-1,-1,0,1,2] [0,0,0,0]
	b := threeSum1(a)
	fmt.Println(b)
}
