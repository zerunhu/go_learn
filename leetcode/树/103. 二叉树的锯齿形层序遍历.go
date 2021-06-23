package main

/*
给定一个二叉树，返回其节点值的锯齿形层序遍历。即先从左往右，再从右往左进行下一层遍历，
以此类推，层与层之间交替进行。

例如：
给定二叉树[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下:
[
  [3],
  [20,9],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	Stack := []*TreeNode{root}
	for level := 0; len(Stack) > 0; level++ {
		var ceng []int
		lenth := len(Stack)
		for i := 0; i < lenth; i++ {
			node := Stack[0]
			Stack = Stack[1:]
			ceng = append(ceng, node.Val)
			if node.Left != nil {
				Stack = append(Stack, node.Left)
			}
			if node.Right != nil {
				Stack = append(Stack, node.Right)
			}
		}
		if level%2 == 0 {
			for i, n := 0, len(ceng); i < n/2; i++ {
				ceng[i], ceng[n-1-i] = ceng[n-1-i], ceng[i]
			}
		}
		res = append(res, ceng)
	}
	return res
}
