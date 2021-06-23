package main

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如，给出
前序遍历 preorder =[3,9,1,5,20,15,7]
中序遍历 inorder = [1,9,5,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9    20
       /  \
1  5  15   7
*/
//前序得到根节点3，中序左边是左节点右边是右节点，根据节点个数可以求得前序1,4为左节点，4,7为右节点
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	var index int
	for i, j := range inorder {
		if j == preorder[0] {
			index = i
			break
		}
	}
	leftInorder := inorder[:index]
	rightInorder := inorder[index+1:]
	leftPreorder := preorder[1 : index+1]
	rightPreorder := preorder[index+1:]
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root
}
