package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file := "E:\\go_learn\\go语言圣经学习笔记\\第一章\\test2.go"
	data, _ := ioutil.ReadFile(file)
	for _, line := range strings.Split(string(data), "\n") {
		fmt.Println(line)
	}
	fmt.Println(data)
	fmt.Println(string(data))
}
