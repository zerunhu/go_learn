package main

import (
	"fmt"
)

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。
在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器。

示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。
在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。

示例 2：
输入：height = [1,1]
输出：1

示例 3：
输入：height = [4,3,2,1,4]
输出：16
*/

//暴力破解超出时间限制giao
func mianji(arr []int) int {
	var mj, bigMj int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				mj = arr[j] * (j - i)
			} else {
				mj = arr[i] * (j - i)
			}
			if mj > bigMj {
				bigMj = mj
			}
		}
	}
	return bigMj
}

//双指针,这里有一个前提就是，在(i,j)这个区间里面，面积是(j-i)*min(ai,aj),如果ai和aj中小的数
//和任意(i,j)区间中的数组成面积都要小于ij的。所以我们每次都移动小的那一边就行了。 O(n)  O(1)
func mianji1(arr []int) int {
	var Bigmj, mj, left, right = 0, 0, 0, len(arr) - 1
	for right >= left {
		if arr[right] > arr[left] {
			mj = arr[left] * (right - left)
			left++
		} else if arr[right] < arr[left] {
			mj = arr[right] * (right - left)
			right--
		} else {
			mj = arr[right] * (right - left)
			right--
			left++
		}
		if mj > Bigmj {
			Bigmj = mj
		}
	}
	return Bigmj
}

func main() {
	testArr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	mj := mianji1(testArr)
	fmt.Println(mj)
}
