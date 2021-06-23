package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//bfs
func Generate(x []int) *TreeNode {
	if len(x) == 0 {
		return nil
	}
	root := &TreeNode{Val: x[0]}
	line := []*TreeNode{root}
	for i := 1; i < len(x); i += 2 {
		node := line[0]
		line = line[1:]
		if x[i] != 0 {
			node.Left = &TreeNode{Val: x[i]}
			line = append(line, node.Left)
		}
		if x[i+1] != 0 && i+1 < len(x) {
			node.Right = &TreeNode{Val: x[i+1]}
			line = append(line, node.Right)
		}
	}
	return root
}

//dfs
//要判断是否有左，右子节点，所有要看数组的长度来判断下一层是否有节点
func helper(nums []int, n int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[n]}
	if len(nums) > 2*n+1 && nums[2*n+1] != 0 {
		root.Left = helper(nums, 2*n+1)
	}
	if len(nums) > 2*n+2 && nums[2*n+2] != 0 {
		root.Right = helper(nums, 2*n+2)
	}
	return root
}
func main() {
	a := []int{1, 2, 0, 4, 5}
	b := helper(a, 0)
	fmt.Println(b)
	fmt.Println(b.Left)
}
