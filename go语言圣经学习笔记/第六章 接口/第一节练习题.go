package main

import (
	"fmt"
	"io"
	"strings"
)

//type ByteCounter struct {
//	w io.Writer
//	n int64
//}

type LimittReader struct {
	R io.Reader
	N int
}

func (l *LimittReader) Read(p []byte) (n int, err error) {
	if l.N < 0 {
		return 0, io.EOF
	}
	if l.N < len(p) {
		p = p[:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= n
	return n, err
}

func test(writer io.Writer, x string) {
	fmt.Println(writer, x)
}
func main() {
	a := strings.NewReader("abcdefg")
	n, _ := a.Read([]byte("abcdefge"))
	//bufio.ScanWords()
	//var c io.LimitReader = io.LimitReader(a, 3)
	fmt.Println(n)
	//fmt.Printf("%T", c)
	//c := &io.LimitedReader{R: a, N: 3}
	//io.Copy(os.Stdout, b)
	//fmt.Println(c)
	//io.Copy(os.Stdout, c)
	//fmt.Println()
	//fmt.Println(c)
	//io.Copy(os.Stdout, c)

	//d := strings.NewReader("abcdefg")
	//n, _ := d.Read([]byte("cd"))
	//io.Copy(os.Stdout, d)
	//fmt.Println(n)
	//c, err := a.Read([]byte("ed"))
	//c := &io.LimitedReader{R: a, N: 3}
	//c.Read()
	//io.Copy(os.Stdout, c)
	//fmt.Println(c)

	//fmt.Printf("%T", a)
	//e := io.Reader()
}
