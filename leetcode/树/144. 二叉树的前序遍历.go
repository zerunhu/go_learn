package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var dg func(node *TreeNode)
	dg = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dg(node.Left)
		dg(node.Right)
	}
	dg(root)
	return res
}
func main() {
	a := TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}
	b := preorderTraversal(&a)
	fmt.Println(b)
}
