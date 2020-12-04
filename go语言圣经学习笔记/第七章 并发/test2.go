package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./test1.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, nerr := io.Copy(os.Stdout, f)
	if nerr != nil {
		fmt.Println(nerr)
		os.Exit(2)
	}

}
