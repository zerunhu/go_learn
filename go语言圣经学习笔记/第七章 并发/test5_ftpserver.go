package main

import (
	"bufio"
	"fmt"
	"net"

	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
}
func Connn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
}
func main() {
	listener, err := net.Listen("tcp", "localhost:800")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		fmt.Println("one connect")
		if err != nil {
			println(err)
			continue
		}
		go Connn(conn)
	}
}
