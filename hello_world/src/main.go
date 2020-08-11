package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello!")
}

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "World!")
}

func world(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "World!")
}

func main() {
	// hello := HelloHandler{}
	// world := WorldHandler{}
	helloHandler := http.HandlerFunc(hello)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// http.Handle("/hello", &hello)
	// http.Handle("/world", &world)
	// http.HandleFunc("/hello", hello)
	http.Handle("/hello", &helloHandler)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}