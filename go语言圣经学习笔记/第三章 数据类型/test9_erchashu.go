package main

import "fmt"

//二叉树排序 通过结构体

func main() {
	t := &tree{value: 5}
	l := [4]int{3, 4, 6, 1}
	for _, i := range l {
		fmt.Println(t)
		t = btree(t, i)
	}
	sort(t)
}

type tree struct {
	value int
	left  *tree
	right *tree
}

func btree(root *tree, node int) *tree {
	if root == nil {
		//var t *tree
		//t = new(tree)
		//t.value = node
		t := &tree{value: node}
		return t
	}
	if node < root.value {
		root.left = btree(root.left, node)
	}
	if node > root.value {
		root.right = btree(root.right, node)
	}
	return root
}
func sort(t *tree) {
	if t.left != nil {
		sort(t.left)
	}
	fmt.Println(t.value)
	if t.right != nil {
		sort(t.right)
	}
}

//    5
//  3   6
//1   4
