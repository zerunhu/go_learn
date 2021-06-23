package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
示例 1：
输入：s = "()"
输出：true
示例2：
输入：s = "()[]{}"
输出：true
示例3：
输入：s = "(]"
输出：false
示例4：
输入：s = "([)]"
输出：false
示例5：
输入：s = "{[]}"
输出：true
左括号入栈，右括号出栈，
*/
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	kh := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var zhan []byte
	for i := 0; i < len(s); i++ {
		if _, ok := kh[s[i]]; !ok {
			zhan = append(zhan, s[i])
		} else {
			// 长度为0代表没有左括号的时候直接右括号，比如"}{"
			if len(zhan) == 0 || zhan[len(zhan)-1] != kh[s[i]] {
				return false
			}
			zhan = zhan[:len(zhan)-1]
		}
	}
	//最后结束了还有剩余比如"{{}"
	return len(zhan) == 0
}
func main() {
	a := "{[]}"
	b := isValid(a)
	fmt.Println(b)
}
