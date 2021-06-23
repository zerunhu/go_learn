package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。


示例 1：
        3
     9    20
        15  7
输入：root = [3,9,20,null,null,15,7]
输出：2
示例 2：

输入：root = [2,null,3,null,4,null,5,null,6]
输出：5

*/
//BFS
//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	m := 1
//	var stack []*TreeNode
//	stack = append(stack, root)
//	for len(stack) > 0 {
//		l := len(stack)
//		for i := 0; i < l; i++ {
//			node := stack[0]
//			stack = stack[1:]
//			if node.Left != nil {
//				stack = append(stack, node.Left)
//			}
//			if node.Right != nil {
//				stack = append(stack, node.Right)
//			}
//			if node.Left == nil && node.Right == nil {
//				return m
//			}
//		}
//		m++
//	}
//	return m
//}

//DFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return minmin(minDepth(root.Left), minDepth(root.Right)) + 1
}
func minmin(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
