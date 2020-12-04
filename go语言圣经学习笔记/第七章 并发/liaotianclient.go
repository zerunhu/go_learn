package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func copymount(dst io.Writer, src io.Reader) {
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
	go copymount(os.Stdout, conn)
	copymount(conn, os.Stdin)
}
