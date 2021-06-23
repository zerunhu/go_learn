package main

import "fmt"

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：
输入：l1 = [], l2 = []
输出：[]
*/
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	l3 := res
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l3.Val = l2.Val
			l2 = l2.Next
		} else {
			l3.Val = l1.Val
			l1 = l1.Next
		}
		l3.Next = &ListNode{}
		l3 = l3.Next
	}
	if l1 != nil {
		l3.Val = l1.Val
		l3.Next = l1.Next
	}
	if l2 != nil {
		l3.Val = l2.Val
		l3.Next = l2.Next
	}
	return res
}
func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}
	b := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}
	c := mergeTwoLists(a, b)
	fmt.Println(c)
	fmt.Println(c.Next)
	fmt.Println(c.Next.Next)
	fmt.Println(c.Next.Next.Next)
	fmt.Println(c.Next.Next.Next.Next)
}
