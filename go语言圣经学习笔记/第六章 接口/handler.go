package main

import (
	"fmt"
	"net/http"
)

type Good map[string]int

//func (g Good) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	panic("implement me")
//}

func (g Good) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/list" {
		for name, _ := range g {
			//writer.Write()
			fmt.Fprintf(writer, "name: %s\n", name)
		}
	} else if request.URL.Path == "/all" {
		for name, price := range g {
			fmt.Fprintf(writer, "name: %s, price: %d\n", name, price)
		}
	} else if request.URL.Path == "/price" {
		name := request.URL.Query().Get("name")
		price, ok := g[name]
		if !ok {
			writer.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(writer, "not have this good: %s\n", name)
			return
		}
		fmt.Fprintf(writer, "name: %s, price: %d\n", name, price)
	} else {
		writer.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(writer, "no such page: %s\n", request.URL)
	}
}
func main() {
	a := Good{"tlbb": 10, "qttlj": 20}
	http.ListenAndServe("localhost:12345", a)
}
