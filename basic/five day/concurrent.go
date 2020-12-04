package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	//say("1")
	go say("2")
	say("3")
	time.Sleep(time.Second)
	fmt.Println("done")
}
