package main

import "fmt"

/*
给定一个包含非负整数的 mxn网格grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

示例 1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。

示例 2：
输入：grid = [[1,2,3],[4,5,6]]
输出：12

如果n=1或m=1，那么最小路径和就是一条线的和，我们先算出这些值，设dp[i][j]为到[i,j]位置的最小路径和，
那么我们可以得出动态转移方程：dp[i][j] += min(dp[i-1][j],dp[i][j-1])。然后计算完事了
O(N^2) O(1)
*/
func minPathSum(grid [][]int) int {
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			if grid[i][j-1] > grid[i-1][j] {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += grid[i][j-1]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(minPathSum(a))
}
