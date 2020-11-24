package main

import "fmt"

func main() {
	var map1 = map[string]string{"f": "F", "a": "A"}
	for i := range map1 {
		fmt.Println(i)
	}
	map2 := map[string]string{"b": "B", "c": "C"}
	i, j := map2["b"]
	fmt.Println(i, j)

	delete(map2, "b")
	fmt.Println(map2)
}
