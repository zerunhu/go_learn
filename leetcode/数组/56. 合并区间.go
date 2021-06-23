package main

import (
	"fmt"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

没啥好说的先排序 在循环判断
O(nlogn) O(n)
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	intervals = fastSort(intervals)
	var res [][]int
	lastInt := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= lastInt[1] {
			if intervals[i][1] > lastInt[1] {
				lastInt = []int{lastInt[0], intervals[i][1]}
			}
		} else {
			res = append(res, lastInt)
			lastInt = intervals[i]
		}
		if i == len(intervals)-1 {
			res = append(res, lastInt)
		}
	}
	return res
}
func fastSort(arr [][]int) [][]int {
	if len(arr) == 0 {
		return arr
	}
	var resArr, leftArr, rightArr [][]int
	keyIndex := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i][0] > keyIndex[0] {
			rightArr = append(rightArr, arr[i])
		} else {
			leftArr = append(leftArr, arr[i])
		}
	}
	resArr = append(fastSort(leftArr), keyIndex)
	resArr = append(resArr, fastSort(rightArr)...)
	return resArr
}
func main() {
	a := [][]int{{1, 4}, {1, 5}}
	b := merge(a)
	fmt.Println(b)
}
