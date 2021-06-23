package main

import "fmt"

/*
翻转一棵二叉树。
输入：
     4
   /   \
  2     7
 / \   / \
1   3 6   9

输出：
     4
   /   \
  7     2
 / \   / \
9   6 3   1
[4,7,2,6,9,1,3]
*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//递归
func bfs1(node *TreeNode) {
	if node == nil {
		return
	}
	bfs1(node.Left)
	bfs1(node.Right)
	node.Left, node.Right = node.Right, node.Left
}
func invertTree(root *TreeNode) *TreeNode {
	bfs1(root)
	return root
}
func qx(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	qx(root.Left)
	qx(root.Right)
}

//深度优先
func bfs2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var Stack []*TreeNode
	Stack = append(Stack, root)
	for len(Stack) > 0 {
		for i := 0; i < len(Stack); i++ {
			node := Stack[len(Stack)-1]
			Stack = Stack[:len(Stack)-1]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				Stack = append(Stack, node.Left)
			}
			if node.Right != nil {
				Stack = append(Stack, node.Right)
			}
		}
	}
	return root
}
func main() {
	a := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6, Left: nil, Right: nil}, Right: &TreeNode{Val: 9, Left: nil, Right: nil}}}
	b := bfs2(a)
	fmt.Println("------------")
	qx(b)
}
