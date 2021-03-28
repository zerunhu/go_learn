package main

import "fmt"

/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
[0,3,5,5,6,7,7,9]   3,4 [0,3,5,6,6,7,7,9]  4,5 [0,3,5,6,7,7,7,9] 5,7[0,3,5,6,7,9,7,9]
*/
/*
数组完成排序后，我们可以放置两个指针 ii 和 jj，其中 ii 是慢指针，而 jj 是快指针。
只要 nums[i] = nums[j]nums[i]=nums[j]，我们就增加 jj 到下一个值

当我们遇到 nums[j] != nums[i]nums[j]=nums[i] 时，跳过重复项的运行已经结束，
因此我们必须把它nums[j]的值复制到 nums[i + 1],因为i-j中间全是重复值，只保留一个i的重复值，
i+1的换成新的不重复，所以先i+1，nums[i]=nums[j]接着我们将再次重复相同的过程
O(n) O(1)
*/
func removeDuplicates(nums []int) int {
	i := 0
	if len(nums) == 0 {
		return 0
	} else {
		for j := 1; j < len(nums); j++ {
			if nums[i] != nums[j] {
				i++
				nums[i] = nums[j]
			}
		}
	}
	return i + 1
}
func main() {
	a := []int{1, 3, 4, 6, 6, 9, 11, 11, 15}
	fmt.Println(removeDuplicates(a))
}
