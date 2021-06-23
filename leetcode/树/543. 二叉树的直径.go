package main

/*
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。
这条路径可能穿过也可能不穿过根结点。
示例 :
给定二叉树
          1
         / \
        2   3
       / \
      4   5
返回3, 它的长度是路径 [4,2,1,3] 或者[5,2,1,3]。
注意：两结点之间的路径长度是以它们之间边的数目表示。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
*/
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var max int
	depth(root, &max)
	return max
}
func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left, max)
	r := depth(root.Right, max)
	*max = max2(l+r, *max)
	return max2(l, r) + 1
}
func max2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
