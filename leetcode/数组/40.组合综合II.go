package main

import (
	"fmt"
	"sort"
)

/*
给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
candidates中的每个数字在每个组合中只能使用一次。
说明：
所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。

示例1:
输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[[1, 7],[1, 2, 5],[2, 6],[1, 1, 6]]
示例2:
输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[[1,2,2],[5]]

和前一题类似
O(n log n)  O(n)
*/

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var a []int
	var b [][]int
	findCombinationSum(candidates, target, 0, a, &b)
	return b
}
func findCombinationSum(arr []int, target int, index int, a []int, b *[][]int) {
	if target < 0 {
		return
	} //10, 1, 2, 7, 6, 1, 5  //1,1,2,5,6,7,10
	if target == 0 {
		c := make([]int, len(a))
		copy(c, a)
		*b = append(*b, c)
		return
	}
	for i := index; i < len(arr); i++ {
		if i > index && arr[i] == arr[i-1] {
			continue
		}
		if target < arr[i] {
			break
		}
		a = append(a, arr[i])
		findCombinationSum(arr, target-arr[i], i+1, a, b)
		a = a[:len(a)-1]
	}
}
func main() {
	a := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	fmt.Println(a)
}
