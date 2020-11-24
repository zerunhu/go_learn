package main

//一个简单的类似于curl命令的函数
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[len(os.Args)-1]
	fmt.Println(url)
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(b))
	fmt.Println(res.Body)
}
