package main

/*
给定一个二叉树，返回其节点值自底向上的层序遍历。即按从叶子节点所在层到根节点所在的层，
逐层从左向右遍历.
例如：
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层序遍历为：

[
  [15,7],
  [9,20],
  [3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
*/
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	line := []*TreeNode{root}
	for len(line) > 0 {
		var ceng []int
		l := len(line)
		for i := 0; i < l; i++ {
			node := line[0]
			ceng = append(ceng, node.Val)
			line = line[1:]
			if node.Left != nil {
				line = append(line, node.Left)
			}
			if node.Right != nil {
				line = append(line, node.Right)
			}
		}
		res = append([][]int{ceng}, res...)
		//这里会造成大量内存消耗，所以直接append最后倒序，用第一个换最后一个依次循环len/2
	}
	return res
}
