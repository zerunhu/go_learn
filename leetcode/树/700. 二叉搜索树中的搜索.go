package main

/*
给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。
返回以该节点为根的子树。 如果节点不存在，则返回 NULL。

例如，
给定二叉搜索树:

        4
       / \
      2   7
     / \
    1   3

和值: 2
你应该返回如下子树:

      2
     / \
    1   3
在上述示例中，如果要找的值是 5，但因为没有节点值为 5，我们应该返回 NULL。
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}
}
func searchBST2(root *TreeNode, val int) *TreeNode {
	for {
		if root == nil || root.Val == val {
			return root
		}
		if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
}
func searchBST3(root *TreeNode, val int) *TreeNode {
	stack := make([]*TreeNode, 0)
	if root.Val == val {
		return root
	} else if root.Val > val && root.Left != nil {
		stack = append(stack, root.Left)
	} else if root.Val < val && root.Right != nil {
		stack = append(stack, root.Right)
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val == val {
			return node
		} else if node.Val > val && node.Left != nil {
			stack = append(stack, node.Left)
		} else if node.Val < val && node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return nil
}
