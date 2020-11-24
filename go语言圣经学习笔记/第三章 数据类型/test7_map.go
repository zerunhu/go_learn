package main

import "fmt"

func main() {
	a := map[string]int{"a": 1, "b": 2}
	if c, ok := a["c"]; !ok { //特殊语法，ok用来判断是否有key,没有!ok就是真
		fmt.Println(1, c)
	} else {
		fmt.Println(2, c)
	}

	var d map[string]map[string]int
	e := map[string]map[string]int{"a": {"a": 1}}
	if d["a"] == nil {
		fmt.Println(11)
		fmt.Println(d)
	}
	fmt.Println(11)
	if e["b"] == nil {
		fmt.Println(11)
		fmt.Println(e)
	}

	f := []string{"a", "b", "a", "c", "d", "c", "d", "f", "d", "a", "b", "s"}
	g := wordfreq(f)
	fmt.Println(g)

	var h map[string]int
	i := make(map[string]int)
	h["a"] = 1 //nil的map不能插入值，map需要初始化，否则报错assignment to entry in nil map
	i["a"] = 1
	fmt.Println(h, i)
}
func wordfreq(s []string) map[string]int {
	a := make(map[string]int)
	for _, j := range s {
		if _, ok := a[j]; !ok {
			a[j] = 1
		} else {
			a[j]++
		}
	}
	return a
}
