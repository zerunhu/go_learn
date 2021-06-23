package main

/*
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差
不超过1，那么它就是一棵平衡二叉树。

给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true


给定二叉树 [1,2,2,3,3,null,null,4,4]
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回false
*/

/*
自顶向下，一边按照前序遍历的顺序一边判断是否平衡
*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func max1(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max1(height(root.Left), height(root.Right)) + 1
}
func abs(a int, b int) bool {
	if a > b {
		return (a - b) <= 1
	} else {
		return (b - a) <= 1
	}
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left), height(root.Right)) && isBalanced(root.Left) && isBalanced(root.Right)
}

/*
自顶向下其实没遍历一次就要求一次height，相当于时间复杂度是N^2
我们自下向顶，用后序遍历的方式，先求左右两个叶子的高，大于2直接就返回false即可。
*/
func isBalanced2(root *TreeNode) bool {
	return houxubianli(root) != -1
}
func houxubianli(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := houxubianli(node.Left)
	if left == -1 {
		return -1
	}
	right := houxubianli(node.Right)
	if right == -1 {
		return -1
	}
	if !abs(left, right) {
		return -1
	}
	return max1(left, right) + 1
}
