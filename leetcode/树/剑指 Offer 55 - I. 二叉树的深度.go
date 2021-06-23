package main

/*
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条
路径，最长路径的长度为树的深度。

给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度3 。
*/
/*
如果我们知道了左子树和右子树的最大深度 ll 和 rr，那么该二叉树的最大深度即为
max(l,r) + 1
而左子树和右子树的最大深度又可以以同样的方式进行计算。
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	length := max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	return length
}
