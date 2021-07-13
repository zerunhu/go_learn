package main

/*
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，
则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，
从而为新的数据值留出空间。

进阶：你是否可以在O(1) 时间复杂度内完成这两种操作？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache
*/
type LRUCache struct {
	size                int
	cap                 int
	HashMap             map[int]*DNodeLink
	LeftHead, RightHead *DNodeLink
}
type DNodeLink struct {
	Key   int //通过这个获取map，就可以吧最后一位删除同时删除这个的map
	Val   int
	Left  *DNodeLink
	Right *DNodeLink
}

func Constructor(capacity int) LRUCache {
	//留前留后为了方便，免得要一直判断nil
	leftHead := &DNodeLink{
		Key:   0,
		Val:   0,
		Right: nil,
	}
	rightHead := &DNodeLink{
		Key:  0,
		Val:  0,
		Left: leftHead,
	}
	leftHead.Right = rightHead
	m := make(map[int]*DNodeLink, capacity)
	return LRUCache{
		size:      0,
		cap:       capacity,
		HashMap:   m,
		LeftHead:  leftHead,
		RightHead: rightHead,
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.HashMap[key]; !ok {
		return -1
	} else {
		//1
		if val.Left != this.LeftHead {
			val.Left.Right = val.Right
			val.Right.Left = val.Left

			val.Left = this.LeftHead
			val.Right = this.LeftHead.Right
			this.LeftHead.Right.Left = val
			this.LeftHead.Right = val
		}
		return val.Val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.HashMap[key]; !ok {
		if this.size+1 > this.cap {
			//1
			va := this.RightHead.Left.Key
			this.RightHead.Left.Left.Right = this.RightHead
			this.RightHead.Left = this.RightHead.Left.Left
			delete(this.HashMap, va)
		} else {
			this.size++
		}
		node := &DNodeLink{
			Key:   key,
			Val:   value,
			Left:  this.LeftHead,
			Right: this.LeftHead.Right,
		}
		this.HashMap[key] = node
		node.Left = this.LeftHead
		node.Right = this.LeftHead.Right
		this.LeftHead.Right.Left = node
		this.LeftHead.Right = node
	} else {
		val.Val = value
		if val.Left != this.LeftHead {
			val.Left.Right = val.Right
			val.Right.Left = val.Left

			val.Left = this.LeftHead
			val.Right = this.LeftHead.Right
			this.LeftHead.Right.Left = val
			this.LeftHead.Right = val
		}
	}
}
