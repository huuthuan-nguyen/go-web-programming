package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP (writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello!")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func (writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("Hanlder called - %T\n", h)
		h.ServeHTTP(writer, request)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler{}

	http.Handle("/hello", log(hello))

	server.ListenAndServe()
}