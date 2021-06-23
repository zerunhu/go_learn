package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for i := 0; i >= 0; i++ {
		if i < n {
			fast = fast.Next
			if fast == nil {
				return head.Next
			}
		} else {
			if fast.Next == nil {
				slow.Next = slow.Next.Next
				break
			}
			slow = slow.Next
			fast = fast.Next
		}
	}
	return head
}
func main() {
	//a := &ListNode{Val: 1, Next: nil}
	a := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	b := removeNthFromEnd(a, 4)
	fmt.Println(b)
}
