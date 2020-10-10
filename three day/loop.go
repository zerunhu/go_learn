package main

import "fmt"

func main() {
	//三元写法
	for sum := 0; sum < 10; sum++ {
		fmt.Println(sum)
	}
	/*
		可省略部分,但是要分号间隔
		sum := 0
		for ; sum < 10; sum++ {
			fmt.Println(sum)
		}
	*/

	sum := 0
	for sum < 10 {
		sum++
		if sum == 5 {
			fmt.Println("a")
			continue
		}
		fmt.Println(sum)
	}
	//为真或者for true
	for {
		sum++
		fmt.Println(sum)
		if sum > 15 {
			break
		}
	}

	//便利数组
	strings := []string{"google", "runoob"}
	for s, j := range strings {
		fmt.Println(s, j)
	}

	/*
		Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
		goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
		但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
	*/
	var a int = 1

	/* 循环 */
asb:
	for a < 5 {
		if a == 3 {
			/* 跳过迭代 */
			a = a + 1
			goto asb
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
