package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//双指针一前一后，然后每次倒序前后指针，然后都后移一位  O(n) O(1)
func reverseList(head *ListNode) *ListNode {
	fast := head //前面的指针是头，后面指针是nil，开局先
	var last *ListNode
	for fast != nil {
		tmp := fast.Next
		fast.Next = last
		last = fast
		fast = tmp
	}
	return last
}

func reverse1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverse1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}

	c := reverse1(a)
	fmt.Println(c)
	fmt.Println(c.Next)
	fmt.Println(c.Next.Next)
	fmt.Println(c.Next.Next.Next)
	fmt.Println(c.Next.Next.Next.Next)
}
