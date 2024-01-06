package main

import (
	"fmt"
	"log"
	"net/http"
)

var count int

func main() {
	c := make(chan int)
	go increment(c)
	http.HandleFunc("/", wrapHandler(handler, c))
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func wrapHandler(handler http.HandlerFunc, c chan int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c <- 1
		handler(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func increment(c <-chan int) {
	for val := range c {
		count += val
	}
}

func counter(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Count %d\n", count)
}
