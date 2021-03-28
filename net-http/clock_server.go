package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(time.Second)
	}
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
		log.Print(conn.RemoteAddr())
		handleConn(conn)
	}
}
