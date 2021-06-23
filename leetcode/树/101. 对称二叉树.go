package main

import (
	"fmt"
)

/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树[1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3

进阶：
你可以运用递归和迭代两种方法解决这个问题吗？


递归
left.right == right.left && left.left == right.right
*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func bfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return bfs(left.Right, right.Left) && bfs(left.Left, right.Right)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return bfs(root.Left, root.Right)
}
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var TreeList []*TreeNode
	TreeList = append(TreeList, root.Left)
	TreeList = append(TreeList, root.Right)
	for len(TreeList) > 0 {
		right := TreeList[len(TreeList)-1]
		TreeList = TreeList[:len(TreeList)-1]
		left := TreeList[len(TreeList)-1]
		TreeList = TreeList[:len(TreeList)-1]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		TreeList = append(TreeList, left.Right)
		TreeList = append(TreeList, right.Left)
		TreeList = append(TreeList, right.Right)
		TreeList = append(TreeList, left.Left)
	}
	return true
}
func main() {
	a := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}}
	fmt.Println(isSymmetric2(a))
}
