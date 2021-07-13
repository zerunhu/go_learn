package main

import "fmt"

/*
给定一个链表，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，
我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没
有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

如果链表中存在环，则返回 true 。 否则，返回 false 。
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
*/
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//方法一通过一个哈希表来存储每一个节点，如果遇到重复节点就是环
func hasCycle(head *ListNode) bool {
	hashMap := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := hashMap[head]; ok {
			return true
		}
		hashMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

//方法二快满指针，快的走两步，慢的走一步，除非是环，否则快的一直在慢的前面
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head.Next
	slow := head
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
func main() {
	a := &ListNode{Val: 1, Next: nil}

	c := hasCycle(a)
	fmt.Println(c)
}
