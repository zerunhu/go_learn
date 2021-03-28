/*
给出两个非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，并且它们的每个节点只能存储一位数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以0开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a := &ListNode{}
	b := a
	jwz := 0
	for l1 != nil || l2 != nil {
		i := l1.Val + l2.Val + jwz
		i, jwz = i%10, i/10
		b.Val = i
		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil && l2 == nil {
			if jwz > 0 {
				b.Next = &ListNode{Val: jwz}
			}
			break
		}
		if l1 == nil {
			l1 = &ListNode{Val: 0}
		}
		if l2 == nil {
			l2 = &ListNode{Val: 0}
		}
		b.Next = &ListNode{}
		b = b.Next
	}
	return a
}

func main() {
	b := &ListNode{3, &ListNode{2, &ListNode{Val: 2}}}
	a := &ListNode{3, &ListNode{2, &ListNode{2, &ListNode{Val: 2}}}}
	x := addTwoNumbers(a, b)
	fmt.Println(x.Val)
	fmt.Println(x.Next.Val)
	fmt.Println(x.Next.Next.Val)
	fmt.Println(x.Next.Next.Next.Val)
}

/*
思路：

*/
