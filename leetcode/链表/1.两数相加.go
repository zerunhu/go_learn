package main

import "fmt"

/*
给出两个非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，并且它们的每个
节点只能存储一位数字。如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以0开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/
type lb struct {
	num  int
	next *lb
}

func add(aLink *lb, bLink *lb) *lb {
	cLink := &lb{}
	c, d := cLink, 0
	for {
		sum := aLink.num + bLink.num + d
		c.num = sum % 10
		d = sum / 10
		aLink = aLink.next
		bLink = bLink.next
		if aLink == nil && bLink == nil {
			if d > 0 {
				c.next = &lb{num: d}
			}
			break
		}
		if aLink == nil {
			aLink = &lb{num: 0}
		}
		if bLink == nil {
			bLink = &lb{num: 0}
		}
		c.next = &lb{}
		c = c.next
	}
	return cLink
}
func main() {
	a := &lb{num: 9, next: &lb{num: 2, next: &lb{num: 3, next: nil}}}
	b := &lb{num: 2, next: &lb{num: 4, next: &lb{num: 9, next: nil}}}
	x := add(a, b)
	fmt.Println(x)
	fmt.Println(x.next)
	fmt.Println(x.next.next)
	fmt.Println(x.next.next.next)
}
