/*
小Q是一个宝石爱好者。这一天，小Q来到了宝石古董店，店家觉得小 Q 是个宝石行家，于是决定和小 Q 玩一个游戏。
游戏是这样的，一共有n块宝石，每块宝石在小Q心中都有其对应的价值。
注意，由于某些宝石质量过于差劲，因此存在只有店家倒贴钱，小 Q 才愿意带走的宝石，即价值可以为负数。
小Q可以免费带走一个连续区间中的宝石，比如区间[1,3]或区间[2,4]中的宝石。请问小Q能带走的最大价值是多少？
*/
package main

import "fmt"

// [1,-3,5,2,-5,2]
func tiaoxuanbaoshi(s []int) int {
	var max int
	l := len(s)
	//for i := l - 1; i >= 0; i-- {
	//	sum += s[l-1]
	//	if max < sum {
	//		max = sum
	//	}
	//}
	maxlist := make([]int, l)
	for i := 0; i < l; i++ {
		if i == 0 {
			maxlist[0] = s[0]
		} else {
			if maxlist[i-1] < 0 {
				maxlist[i] = s[i]
			} else {
				maxlist[i] = maxlist[i-1] + s[i]
			}
		}
		if max < maxlist[i] {
			max = maxlist[i]
		}
	}
	return max
}
func main() {
	a := []int{1, 2, 5, 2, -1, 2, -5, -3, 9}
	b := tiaoxuanbaoshi(a)
	fmt.Println(b)
}
