package main

import "fmt"

/*
给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
k是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]

输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

输入：head = [1], k = 1
输出：[1]
*/
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/*
先吧链表分成k个一组，然后反转，中间的连接非常重要，具体是每次要留前一个链表组的最后一位值，然后next就
是当前链表组倒叙之后的第一个，第一组没有pre可以在前面在家一个head当作pre就不用做特殊判断了
一直每k次循环倒叙一组一边循环一边倒叙，和其他人先循环完判断大于k在倒叙不同。如果在k次循环中nil了，
说明最后一组不够k个，在吧最后一组倒叙回去。
O(N) O(1)
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	fast := head
	toutou := &ListNode{Val: 0, Next: head}
	pre := toutou //pre是当前链表组前一个的最后一个值，用来连接当前链表组反转之后的第一个值
	for fast != nil {
		var slow *ListNode
		start := fast //start是当前链表组的第一个值，反转之后就是最后一个值了，就是下一轮链表组的pre
		for i := 0; i < k; i++ {
			if fast == nil {
				fast, slow = slow, fast
				for fast != nil {
					tmp := fast.Next
					fast.Next = slow
					slow = fast
					fast = tmp
				}
				break
			}
			tmp := fast.Next
			fast.Next = slow
			slow = fast
			fast = tmp
		}
		pre.Next = slow //反转之后slow是前一个链表组的最后一个，fast是后一个链表组的第一个
		pre = start
	}
	return toutou.Next
}
func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}}}}}}

	c := reverseKGroup(a, 3)
	fmt.Println(c)
	fmt.Println(c.Next)
	fmt.Println(c.Next.Next)
	fmt.Println(c.Next.Next.Next)
	fmt.Println(c.Next.Next.Next.Next)
	fmt.Println(c.Next.Next.Next.Next.Next)
	fmt.Println(c.Next.Next.Next.Next.Next.Next)
}
