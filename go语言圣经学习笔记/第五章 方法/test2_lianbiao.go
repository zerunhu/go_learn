package main

import "fmt"

type IntList struct {
	value int
	next  *IntList
}

func (i *IntList) geneList(s []int) {
	cur := i
	for _, j := range s {
		next := &IntList{j, nil}
		cur.next = next
		cur = next
	}
}
func geneList(i *IntList, s []int) *IntList {
	cur := i
	for _, j := range s {
		next := &IntList{j, nil}
		cur.next = next
		cur = next
	}
	return i
}
func (i *IntList) sum() int {
	if i == nil {
		return 0
	}
	return i.value + i.next.sum()
}
func main() {
	a := IntList{1, nil}
	b := []int{1, 2, 3}
	//a.geneList(b)
	c := geneList(&a, b)
	d := a.sum()
	fmt.Println(c.next.next, d)
}
