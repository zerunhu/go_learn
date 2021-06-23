package main

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]
*/
import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func swapPairs(head *ListNode) *ListNode {
	last := &ListNode{Next: head}
	res := last
	for last.Next != nil && last.Next.Next != nil {
		node1 := last.Next
		node2 := last.Next.Next
		last.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		last = node1
	}
	return res.Next
}
func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}
	//b := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}
	c := swapPairs(a)
	fmt.Println(c)
	fmt.Println(c.Next)
	fmt.Println(c.Next.Next)
	fmt.Println(c.Next.Next.Next)
}
