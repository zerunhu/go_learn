package main

import "fmt"

type book struct {
	title  string
	author string
	id     int
}

func main() {
	var a = book{title: "a", author: "1", id: 1}
	fmt.Print(a)
	fmt.Println(a.title)

	var b book
	c := printbook(b)
	fmt.Print(c)
}
func printbook(book book) book {
	book.id = 1
	book.author = "1"
	book.title = "1"
	return book
}
