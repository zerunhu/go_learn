package main

import (
	"fmt"
	"sort"
)

/*
给定一个无重复元素的数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
candidates中的数字可以无限制重复被选取。

说明：
所有数字（包括target）都是正整数。
解集不能包含重复的组合。

示例：
输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]

示例2：
输入：candidates = [2,3,5], target = 8,
所求解集为：
[
[2,2,2,2],
[2,3,3],
[3,5]
]
*/

/*
讲实话 我看了分析之后能懂，但是让我自己在做 我肯定做不出来 -~-
O(n log n) O(n)
这里数组的本质是多个类型接口的指针，所以记得copy在append
golang切片不会产生新地址，python会产生新地址
*/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var c []int
	if len(candidates) == 0 {
		return res
	}
	sort.Ints(candidates)
	findcombinationSum(candidates, target, 0, c, res)
	return res
}
func findcombinationSum(candidates []int, target int, index int, c []int, res [][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		b := make([]int, len(c))
		copy(b, c)
		res = append(res, b)
		return
	}
	for i := index; i < len(candidates); i++ {
		if target < candidates[i] {
			break
		}
		c = append(c, candidates[i])
		findcombinationSum(candidates, target-candidates[i], i, c, res)
		c = c[:len(c)-1]
	}
}
func main() {
	a := combinationSum([]int{2, 3, 5}, 8)
	fmt.Println(a)
}
