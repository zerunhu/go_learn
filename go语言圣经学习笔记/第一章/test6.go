package main

//一个简单的类似于curl命令的函数
import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v\n", err)
		os.Exit(1)
	}
	//b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "read error: %v\n", err)
	//	os.Exit(1)
	//}
	ch <- string(res.Status)
}

func main() {
	ch := make(chan string)
	go fetch("www.baidu.com", ch)
	go fetch("www.zhihu.com", ch)
	a := <-ch
	b := <-ch
	fmt.Println(a, b)
}
