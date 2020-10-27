package main

import "fmt"

type errstr struct {
	shang int
	xia   int
}

func (cf errstr) err() string {
	aa := "fasheng error shang: %d , xia: %d"
	return fmt.Sprintf(aa, cf.shang, cf.xia)
}

func chufa(shang int, xia int) (result int, x string) {
	if xia == 0 {
		errmsg := errstr{shang: shang, xia: 0}
		x = errmsg.err()
		return
	} else {
		result = shang / xia
		return
	}
}

func main() {
	a, b := chufa(100, 10)
	fmt.Println(a, b)
	c, d := chufa(100, 0)
	fmt.Println(c, d)
}
