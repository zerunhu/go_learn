package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func wait(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	go wait(100 * time.Millisecond)
	n := 45
	f := fib(n)
	fmt.Println("111111", f)
}
