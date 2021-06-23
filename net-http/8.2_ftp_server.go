package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func Shellout1(command string) string {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String()
	}
	return stdout.String()
}
func fileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return true
}
func handleConn1(c net.Conn) {
	defer log.Printf("a client duankai: %s\n", c.RemoteAddr())
	defer c.Close()
	input := bufio.NewScanner(c)
	fmt.Println(-1)
	for input.Scan() {
		fmt.Println(00)
		cmd := input.Text()
		if cmd == "quit" || cmd == "break" {
			c.Write([]byte("client quit, client duankai\n"))
			return
		}
		if strings.HasPrefix(cmd, "cd") {
			path := cmd[3:]
			os.Chdir(path)
			continue
		}
		if strings.HasPrefix(cmd, "upload") {
			fmt.Println(111111, cmd)
			cmdSplit := strings.Split(cmd, ",")
			fileName := cmdSplit[0][7:]
			if fileExist(fileName) {
				fmt.Fprintf(c, "%s exist, upload failed", fileName)
			}
			f, err1 := os.Create(fileName)
			if err1 != nil {
				log.Fatal(err1)
			}
			n, err2 := io.WriteString(f, cmdSplit[1])
			if err2 != nil {
				log.Fatal(err2)
			}
			f.Close()
			fmt.Printf("写入文件%s->%d个字节\n", fileName, n)
			continue
		}
		fmt.Println(11)
		cmdResult := Shellout1(cmd)
		fmt.Println(22)
		fmt.Printf("client: %s exec cmd: %s\n%s", c.RemoteAddr(), cmd, cmdResult)
		c.Write([]byte(cmdResult))
		fmt.Println(33)
	}
	fmt.Println(44)
}
func main() {
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Print(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		log.Printf("a client connect: %s", conn.RemoteAddr())
		go handleConn1(conn)
	}
}
