package main

import "fmt"

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在
大于或等于 x 的节点之前。
你应当 保留 两个分区中每个节点的初始相对位置。

输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

输入：head = [2,1], x = 2
输出：[1,2]

分两个链表，最后合并链表
*/
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func partition(head *ListNode, x int) *ListNode {
	smallList := &ListNode{}
	sHead := smallList
	bigList := &ListNode{}
	bHead := bigList
	for head != nil {
		if head.Val < x {
			smallList.Next = head
			smallList = smallList.Next
			head = head.Next
		} else {
			bigList.Next = head
			bigList = bigList.Next
			head = head.Next
		}
	}
	smallList.Next = bHead.Next
	bigList.Next = nil
	return sHead.Next
}
func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: nil}}}}}}
	c := partition(a, 3)
	fmt.Println(c)
	fmt.Println(c.Next)
	fmt.Println(c.Next.Next)
	fmt.Println(c.Next.Next.Next)
	fmt.Println(c.Next.Next.Next.Next)
	fmt.Println(c.Next.Next.Next.Next.Next)
}
