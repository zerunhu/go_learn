package main

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true

*/
func isPalindrome(head *ListNode) bool {
	nHead := &ListNode{Next: head}
	fast, slow := nHead, nHead
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	hList := slow.Next

	right := hList
	var left *ListNode
	left = nil
	for right != nil {
		tmp := right.Next
		right.Next = left
		left = right
		right = tmp
	}
	for left != nil {
		if left.Val != head.Val {
			return false
		}
		left = left.Next
		head = head.Next
	}
	return true
}
