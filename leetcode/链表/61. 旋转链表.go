package main

import "fmt"

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
输入：head = [0,1,2], k = 4
输出：[2,0,1]

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]
*/
//成环  然后循环就行 比较简单，注意不是直接往后移动就行，这里出错过
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	nList := head
	start := head
	n := 1
	for nList.Next != nil {
		nList = nList.Next
		n++
	}
	k = n - (k % n)
	end := nList
	nList.Next = start
	for i := 0; i < k; i++ {
		start = start.Next
		end = end.Next
	}
	end.Next = nil
	return start
}
func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	c := rotateRight(a, 6)
	fmt.Println(c)
	fmt.Println(c.Next)
	fmt.Println(c.Next.Next)
	fmt.Println(c.Next.Next.Next)
	fmt.Println(c.Next.Next.Next.Next)
}
