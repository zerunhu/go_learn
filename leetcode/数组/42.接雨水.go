package main

/*
给定n个非负整数表示每个宽度为1的柱子的高度图，计算按此排列的柱子,下雨之后能接多少雨水。
示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

*/

/*就是算两个高柱子中间的空间，双指针一左一右，定义两个标记位，算是单边最高的位置，左右指针小的那个
移动，如果移动之后的值大于标记位置说明已经形成了一个小水池了，那样的话如果小于标记位置就是在水池
中间，累加就完事了，所以一波循环知道双指针相遇结束
O(N) O(1)
下面注释的是我没看答案自己写的，快慢指针都从左开始，当遇到下坡的时候相当于形成了一个水池的左边了。
标记一下说明水池开始，然后慢指针不动，快指针一直到右边大于水池左边的高度，组成一个水池，然后计算
里面的值。有一个边际情况就是右边可能一直不大于左边的高数组就结束了，其实就是左高右低形成的一个水池
然后其实就是一个倒着的前面讲的方法，直接数组倒序在调用自己完事。
这个时间复杂度最差O(n^3) 最好O(n)
*/
import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	leftBigNum, rightBigNum, sum := height[left], height[right], 0
	for left < right {
		if height[left] > height[right] {
			right--
			if height[right] <= rightBigNum {
				sum += rightBigNum - height[right]
			} else {
				rightBigNum = height[right]
			}
		} else {
			left++
			if height[left] <= leftBigNum {
				sum += leftBigNum - height[left]
			} else {
				leftBigNum = height[left]
			}
		}
	}
	return sum
}

//func trap(height []int) int {
//	//var arr []int
//	bj, low, sum, lastSum := 0, 0, 0, 0
//	for fast := 1; fast < len(height); fast++ {
//		if bj == 0 {
//			if height[fast] >= height[low] {
//				low++
//			} else {
//				bj = 1
//			}
//		} else { //{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
//			if height[fast] >= height[low] {
//				num := jsmj(height[low : fast+1])
//				bj = 0
//				low = fast
//				sum += num
//			}
//			if fast == len(height)-1 && height[fast] < height[low] {
//				arr := height[low:]
//				length := len(arr)
//				for i := 0; i < length/2; i++ {
//					temp := arr[length-1-i]
//					arr[length-1-i] = arr[i]
//					arr[i] = temp
//				}
//				lastSum = trap(arr)
//			}
//		}
//	}
//	return sum + lastSum
//}
//func jsmj(arr []int) int {
//	//fmt.Println(arr)
//	var height int
//	if arr[0] >= arr[len(arr)-1] {
//		height = arr[len(arr)-1]
//	} else {
//		height = arr[0]
//	}
//	width := len(arr) - 2
//	bigMj := height * width
//	//println(height, width, bigMj)
//	for i := 1; i < len(arr)-1; i++ {
//		if arr[i] < height {
//			bigMj -= arr[i]
//		} else {
//			bigMj -= height
//		}
//	}
//	return bigMj
//}
func main() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	b := trap(a)
	fmt.Println(b)
}
