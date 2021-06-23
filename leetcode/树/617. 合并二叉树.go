package main

/*
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠
那么将他们的值相加作为节点合并后的新值，否则不为NULL 的节点将直接作为新二叉树的节点。
注意: 合并必须从两个树的根节点开始。

输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7

[3,1,5,null,4,null,7]
预期结果
[3,4,5,5,4,null,7]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-binary-trees

*/
func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees1(root1.Left, root2.Left)
	root1.Right = mergeTrees1(root1.Right, root2.Right)
	return root1
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root := &TreeNode{Val: root1.Val + root2.Val}
	line, line1, line2 := []*TreeNode{root}, []*TreeNode{root1}, []*TreeNode{root2}
	for len(line1) > 0 && len(line2) > 0 {
		node := line[0]
		line = line[1:]
		node1 := line1[0]
		line1 = line1[1:]
		node2 := line2[0]
		line2 = line2[1:]
		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if left1 != nil || left2 != nil {
			if left1 != nil && left2 != nil {
				left := &TreeNode{Val: left1.Val + left2.Val}
				node.Left = left
				line = append(line, left)
				line1 = append(line1, left1)
				line2 = append(line2, left2)
			} else if left1 != nil {
				node.Left = left1
			} else {
				node.Left = left2
			}
		}
		if right1 != nil || right2 != nil {
			if right1 != nil && right2 != nil {
				right := &TreeNode{Val: right1.Val + right2.Val}
				node.Right = right
				line = append(line, right)
				line1 = append(line1, right1)
				line2 = append(line2, right2)
			} else if right1 != nil {
				node.Right = right1
			} else {
				node.Right = right2
			}
		}
	}
	return root
}
