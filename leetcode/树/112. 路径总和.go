package main

import "fmt"

/*
给你二叉树的根节点root 和一个表示目标和的整数targetSum ，
判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。

示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
示例 2：
输入：root = [1,2,3], targetSum = 5
输出：false

示例 3：
输入：root = [1,2], targetSum = 0
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum
*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//bfs
//func hasPathSum(root *TreeNode, targetSum int) bool {
//	if root == nil {
//		return false
//	}
//	line1 := []int{root.Val}
//	line2 := []*TreeNode{root}
//	for len(line2) > 0 {
//		l := len(line2)
//		for i := 0; i < l; i++ { //层序遍历
//			node := line2[0]
//			value := line1[0]
//			line1 = line1[1:]
//			line2 = line2[1:]
//			if node.Left == nil && node.Right == nil {
//				if value == targetSum { //value跟着node走，就是从上到下这个node的路径和
//					return true
//				}
//			}
//			if node.Left != nil {
//				line1 = append(line1, value+node.Left.Val)
//				line2 = append(line2, node.Left)
//			}
//			if node.Right != nil {
//				line1 = append(line1, value+node.Right.Val)
//				line2 = append(line2, node.Right)
//			}
//		}
//	}
//	return false
//}

//dfs
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	a := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}
	fmt.Print(hasPathSum(&a, 2))
}
