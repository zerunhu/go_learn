package main

import (
	"fmt"
	"sort"
)

func lengthOfLongestSubstring(s string) int {
	a := []int{}
	for i, j := range s {
		b := map[int32]int{}
		b[j] = 0
		for _, y := range s[i+1:] {
			if _, ok := b[y]; !ok {
				b[y] = 0
			} else {
				break
			}
		}
		a = append(a, len(b))
	}
	sort.Sort(sort.IntSlice(a))
	if len(a) == 0 {
		return 0
	}
	return a[len(a)-1]
}
func lengthOfLongestSubstring1(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		if ans < rk-i+1 {
			ans = rk - i + 1
		}
	}
	return ans
}
func lengthOfLongestSubstring2(s string) int {
	b := len(s)
	nas := 0
	for i := 0; i < b; i++ {
		a := map[byte]int{}
		//c := s[i:]
		left := -1 + i
		for left+1 < b && a[s[left+1]] != 1 {
			a[s[left+1]]++
			left++
		}
		if left+1-i > nas {
			nas = left + 1 - i
		}
	}
	return nas
}

func main() {
	s := ""
	x := lengthOfLongestSubstring2(s)
	fmt.Println(x)
}
