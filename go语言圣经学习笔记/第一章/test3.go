package main

import (
	"bufio"
	"fmt"
	"os"
)

func pfile(f *os.File) {
	input := bufio.NewScanner(f)
	fmt.Println(input)
}

func main() {
	file := "E:\\go_learn\\go语言圣经学习笔记\\第一章\\test2.go"
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	} else {
		input := bufio.NewScanner(f)
		for input.Scan() {
			fmt.Println(input.Text())
		}
	}
	fmt.Println(f)
	var ff int = 20
	fmt.Println(ff)
	fmt.Println(&ff)
	b := [1]int{ff}
	fmt.Println(&b)
	//fmt.Println(*"0xc0000a2130")
}
