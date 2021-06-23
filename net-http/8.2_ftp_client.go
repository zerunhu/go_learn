package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func ioInput(conn net.Conn) {
	fmt.Print("Please input cmd: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "input error:", err)
	}
	cmd := scanner.Text()
	if strings.HasPrefix(cmd, "upload") {
		fileName := strings.TrimSpace(cmd[7:])
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		cmd = cmd + "," + string(content)
		fmt.Println(cmd)
	}
	fmt.Fprintln(conn, cmd)
	fmt.Println(2)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	done := make(chan struct{})
	//go mustCopy1(os.Stdout, conn, done)
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println(1111)
		done <- struct{}{}
		fmt.Println(2222)
	}()
	for {
		fmt.Println(1)
		ioInput(conn)
		fmt.Println(3)
		<-done
		fmt.Println(4)
	}
}
func mustCopy1(dst io.Writer, src io.Reader, done chan struct{}) {
	if _, err := io.Copy(dst, src); err != nil {
		done <- struct{}{}
		log.Fatal(err)
	}
}
func mustCopy2(dst io.Writer, src io.Reader) {
	fmt.Println("success connect ftp")
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
