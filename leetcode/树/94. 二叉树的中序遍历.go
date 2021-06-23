package main

/*
前序  先访问节点在访问左叶子在访问右叶子
中序  先访问左叶子在访问节点在访问右叶子
后续  先访问左叶子在访问右叶子在访问节点
*/

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func inorderTraversal1(root *TreeNode) (res []int) {
	var dg func(node *TreeNode) //为了利用外部的变量而使用闭包
	dg = func(node *TreeNode) {
		if node == nil {
			return
		}
		dg(node.Left)
		res = append(res, node.Val)
		dg(node.Right)
		return
	}
	dg(root)
	return
}
func inorderTraversal(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

func main() {
	a := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}
	b := inorderTraversal(&a)
	fmt.Println(b)
}
