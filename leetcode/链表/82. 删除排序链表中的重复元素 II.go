package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/*
存在一个按升序排列的链表,给你这个链表的头节点head,请你删除链表中所有存在数字重复情况的节点，
只保留原始链表中没有重复出现的数字。
返回同样按升序排列的结果链表。
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
输入：head = [1,1,1,2,3]
输出：[2,3]

没啥好说的 遇到相等 不能都删因为要和后面继续比较是否相等，设置一个key来决定前面一个是否需要删，当前后
不相等的时候，再来吧前面一直保留的想等值删了。然后key初始。需要一个头留着因为head头可能被删除
*/
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head
	slow := &ListNode{Next: head}
	tou := slow
	dupKey := 0
	for fast.Next != nil {
		if fast.Val == fast.Next.Val {
			slow.Next = fast.Next
			fast = fast.Next
			dupKey = 1
		} else {
			if dupKey == 1 {
				slow.Next = fast.Next
				fast = fast.Next
				dupKey = 0
			} else {
				slow = slow.Next
				fast = fast.Next
			}
		}
	}
	if dupKey == 1 {
		slow.Next = fast.Next
	}
	return tou.Next
}
func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}}}
	c := deleteDuplicates2(a)
	fmt.Println(c)
	fmt.Println(c.Next)
	fmt.Println(c.Next.Next)
}
