package main

/*
给你二叉树的根节点root和一个整数目标和Sum,找出所有从根节点到叶子节点路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。

示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]

示例 2：
输入：root = [1,2,3], targetSum = 5
输出：[]

示例 3：
输入：root = [1,2], targetSum = 0
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-ii
*/
//dfs
//func pathSum(root *TreeNode, targetSum int) [][]int {
//	var dfs func(node *TreeNode, sum int, sumList []int)
//	var res [][]int
//	dfs = func(node *TreeNode, sum int, sumList []int) {
//		if node == nil {
//			return
//		}
//		sumList = append(sumList, node.Val)
//		if node.Left == nil && node.Right == nil {
//			if sum == node.Val {
//				res = append(res, sumList)
//				return
//			}
//			return
//		}
//		//这里为了防止引用同一个数组进行左右子节点操作，最后对结果产生错误
//		copysumList := make([]int, len(sumList))
//		copy(copysumList, sumList)
//		dfs(node.Left, sum-node.Val, copysumList)
//		dfs(node.Right, sum-node.Val, copysumList)
//	}
//	dfs(root, targetSum, []int{})
//	return res
//}

//这个是官方的方法，不需要那么多数组，用一个数组每次函数执行完回去就pop一下，减少了很多内存空间，
func pathSum(root *TreeNode, sum int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, sum)
	return
}
