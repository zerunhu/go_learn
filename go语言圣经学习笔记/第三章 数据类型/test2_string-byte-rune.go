package main

import (
	"fmt"
)

func main() {
	a := "abc" //字符串实际上是以byte[]的形式存储在底层的，byte就是uint8
	b := []byte(a)
	b[1] = 99
	c := string(b)
	d := "钉钉"
	fmt.Println(a, b, c, []rune(c))
	fmt.Println([]rune(d), []byte(d))
	//h := "ilove中国"
	//fmt.Println(a, h, len(h)) //golang中string底层是通过byte数组实现的。
	//// 中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
	//i := []rune(h)
	//fmt.Println(i, len(i))
	///*
	//	golang中还有一个byte数据类型与rune相似，它们都是用来表示字符类型的变量类型。它们的不同在于：
	//	byte 等同于int8，常用来处理ascii字符
	//	rune 等同于int32,常用来处理unicode或utf-8字符
	//*/
}
