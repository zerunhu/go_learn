package main

import "fmt"

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]
输入：head = [5], left = 1, right = 1
输出：[5]
没啥特别好说的，链表都那样，一边循环跳过前left的，然后开始和206一样双指针反转，最后拼起来
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	pre := dummyNode
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	end := pre.Next
	slow := pre.Next
	fast := pre.Next.Next
	for i := 0; i < right-left; i++ {
		tmp := fast.Next
		fast.Next = slow
		slow = fast
		fast = tmp
	}
	pre.Next = slow
	end.Next = fast
	return dummyNode.Next
}

var a *ListNode

func reverse(head *ListNode, right int) *ListNode {
	if right == 1 && head != nil && head.Next != nil {
		a = head.Next
		return head
	}
	last := reverse(head.Next, right-1)
	head.Next.Next = head
	head.Next = a
	return last
}
func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}

	c := reverse(a, 3)
	fmt.Println(c)
	fmt.Println(c.Next)
	fmt.Println(c.Next.Next)
	fmt.Println(c.Next.Next.Next)
}
