package main

import "fmt"

/*
给你一个字符串 s，找到 s 中最长的回文子串。
示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

示例 3：
输入：s = "a"
输出："a"

这里用dp，dp[i,j]代表s[i:j]是否是回文串，那么dp[i,j] = dp[i+1,j-1] && i =j ,确定了状态转移方程
那么就要判断左下角的值来确定现在的值，可以两次循环获取，先获取l从小到大，2->len,在获取开头的值i，然后
根据i和l来获取j，为什么要这样是因为，通过类似于一个方格的来说实现建模，从左上角到右下角的一个斜线，代表
i=j,即一个数字，那么肯定是true，继续就是长度为2的，即i+1=j，即第一条斜线的向右平移一格的斜线，一直到
最后一个i=0.j=len的时候，那么就是l可以取最长的时候，然后循环一次都判断一下是不是最长的值，是就换
O(N^2)  O(N^2)
*/

func longestPalindrome(s string) string {
	len := len(s)
	left, right := 0, 0
	dp := make([][]bool, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]bool, len)
		dp[i][i] = true
	}
	for l := 2; l <= len; l++ {
		for i := 0; i < len; i++ {
			j := i + l - 1
			if j >= len {
				break
			}
			if s[i] == s[j] {
				if l < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] && l > (right-left) {
				left, right = i, j
			}
		}
	}
	return s[left : right+1]
}
func main() {
	a := "aa"
	fmt.Println(longestPalindrome(a))
}
