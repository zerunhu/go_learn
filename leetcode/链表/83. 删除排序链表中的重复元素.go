package main

import "fmt"

/*
存在一个按升序排列的链表,给你这个链表的头节点head，请你删除所有重复的元素,使每个元素,只出现一次
返回同样按升序排列的结果链表。
*/
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil {
		if fast.Val == slow.Val {
			slow.Next = fast.Next
			fast = fast.Next
		} else {
			slow = slow.Next
			fast = fast.Next
		}
	}
	return head
}
func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}}}
	c := deleteDuplicates(a)
	fmt.Println(c)
	fmt.Println(c.Next)
	fmt.Println(c.Next.Next)
}
