package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func Conn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05 MST Mon\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
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
		go Conn(conn)
	}
}
