package main

import (
	"fmt"
)

/*
给你一个整数数组nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：
输入：nums = [0]
输出：[[],[0]]


感觉像dp不知道是不是，
从第一个开始,dp[n] = dp[n-1]+dp[n-1]的每个子集加上一个n
[]-->[]+[1]-->[][1]+[2][1,2]-->[][1][2][1,2]+[3][1,3][2,3][1,2,3]
*/
func subsets(nums []int) [][]int {
	res := [][]int{{}}
	for _, i := range nums {
		l := len(res)
		for j := 0; j < l; j++ {
			temp := make([]int, len(res[j]))
			copy(temp, res[j])
			tempL := append(temp, i)
			res = append(res, tempL)
		}
	}
	return res
}
func main() {
	a := []int{9, 0, 3, 5, 7}
	b := subsets(a)
	fmt.Println(b)
	fmt.Printf("%p-%p-%p", b[7], b[15], b[23])
}
