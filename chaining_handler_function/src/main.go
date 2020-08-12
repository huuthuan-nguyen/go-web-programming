package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// handler func
func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello!")
}

// chaining handler
func log(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		handler(writer, request)
	}
}

func main() {
	h := http.HandlerFunc(hello)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", log(h))
	server.ListenAndServe()
}