package main

import (
	"fmt"
	"math"
	"strings"
)

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
示例 2：
输入：s = "a", t = "a"
输出："a"

来源：力扣（LeetCode）  math.MinInt32
链接：https://leetcode-cn.com/problems/minimum-window-substring
*/
func minWindow(s string, t string) string {
	left, right := 0, 0
	min := math.MaxInt8
	str := ""
	hashMap := make(map[byte]int, len(t))
	for ; right < len(s); right++ {
		//fmt.Println(left, right, hashMap)
		if strings.IndexByte(t, s[right]) > -1 {
			hashMap[s[right]] += 1
			if len(hashMap) == len(t) {
				length := right - left + 1
				if length < min {
					min = length
					str = s[left : right+1]
				}
				for left < right {
					if _, ok := hashMap[s[left]]; ok {
						hashMap[s[left]] -= 1
						if hashMap[s[left]] == 0 {
							delete(hashMap, s[left])
						}
					}
					left++
					if len(hashMap) != len(t) {
						break
					} else {
						length = right - left + 1
						if length < min {
							min = length
							str = s[left : right+1]
						}
					}
				}
			}
		}
	}
	return str
}

func main() {
	s := "aa"
	t := "aa"
	a := minWindow(s, t)
	fmt.Println(a)
}
