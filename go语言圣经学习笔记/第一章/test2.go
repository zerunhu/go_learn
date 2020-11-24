package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//counts[input.Text()]++
		fmt.Println("1")
	}
	//a := map[string]int{"aa": 1, "cc": 2}
	//b := a["ab"]
	//fmt.Println(counts)
}
