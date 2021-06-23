package main

/*
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入:[1,2,3,null,5,null,4]
输出:[1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
*/

//bfs
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	line := []*TreeNode{root}
	res := []int{root.Val}
	for len(line) > 0 {
		l := len(line)
		for i := 0; i < l; i++ {
			node := line[0]
			line = line[1:]
			if node.Left != nil {
				line = append(line, node.Left)
			}
			if node.Right != nil {
				line = append(line, node.Right)
			}
		}
		if len(line) > 0 {
			res = append(res, line[len(line)-1].Val)
		}
	}
	return res
}

//dfs
func rightSideView1(root *TreeNode) []int {
	var res []int
	var test func(node *TreeNode, depth int)
	test = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		//数组存值第一个即第一层
		if len(res) == depth {
			res = append(res, node.Val)
		}
		test(node.Right, depth+1)
		test(node.Left, depth+1)
	}
	test(root, 0)
	return res
}
