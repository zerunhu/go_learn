package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.Minute * 2
	deadline := time.Now().Add(timeout)
	aa := time.Now().Before(deadline)
	fmt.Println(timeout)
	fmt.Println(deadline)
	fmt.Println(aa)
}
