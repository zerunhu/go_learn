package main

import "fmt"

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
示例：
二叉树：[3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：
[
  [3],
  [9,20],
  [15,7]
]

*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var Stack []*TreeNode
	Stack = append(Stack, root)
	for len(Stack) > 0 {
		var ceng []int
		lenth := len(Stack)
		for i := 0; i < lenth; i++ {
			node := Stack[0]
			Stack = Stack[1:]
			ceng = append(ceng, node.Val)
			if node.Left != nil {
				Stack = append(Stack, node.Left)
			}
			if node.Right != nil {
				Stack = append(Stack, node.Right)
			}
		}
		res = append(res, ceng)
	}
	return res
}
func main() {
	a := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6, Left: nil, Right: nil}, Right: &TreeNode{Val: 9, Left: nil, Right: nil}}}
	b := levelOrder(a)
	fmt.Println(b)

}
