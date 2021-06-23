package main

import (
	"fmt"
	"math"
)

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例1:
输入:
    2
   / \
  1   3
输出: true

示例2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
根节点的值为 5 ，但是其右子节点值为 4 。
*/

/*
递归：左小右大很简单
循环：中序遍历，然后递增
*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isValidBST(root *TreeNode) bool {
	var bfs func(node *TreeNode, left int, right int) bool
	bfs = func(node *TreeNode, left int, right int) bool {
		if node == nil {
			return true
		}
		if node.Val <= left || node.Val >= right {
			return false
		}
		return bfs(node.Left, left, node.Val) && bfs(node.Right, node.Val, right)
	}
	return bfs(root, math.MinInt32, math.MaxInt32)
}
func main() {
	a := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6, Left: nil, Right: nil}, Right: &TreeNode{Val: 9, Left: nil, Right: nil}}}
	b := isValidBST(a)
	fmt.Println(b)
}
