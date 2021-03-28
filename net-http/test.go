package main

import (
	"fmt"
	"time"
)

func sninner(delay time.Duration) {
	for {
		for _, j := range `-\|/` {
			fmt.Printf("\r%c", j)
			time.Sleep(delay)
		}
	}
}
func fblq(n int) int {
	if n < 2 {
		return n
	} else {
		return fblq(n-1) + fblq(n-2)
	}
}

func main() {
	go sninner(100 * time.Millisecond)
	f := fblq(45)
	fmt.Println(f)
}
