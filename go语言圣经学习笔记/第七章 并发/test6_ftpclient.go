package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func printStdout1(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	conn, err := net.Dial("tcp", "localhost:800")
	if err != nil {
		println(err)
	}
	defer conn.Close()
	go printStdout1(os.Stdout, conn)
	printStdout1(conn, os.Stdin)
}
