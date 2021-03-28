package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Good1 map[string]int

func (db Good1) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %d\n", item, price)
	}
}

func (db Good1) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "name: %s, price: %d\n", item, price)
}

func (db Good1) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	price := req.URL.Query().Get("price")
	price1, _ := strconv.Atoi(price)
	db[item] = price1
	fmt.Fprintf(w, "%s\n", price)
}
func main() {
	db := Good1{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
