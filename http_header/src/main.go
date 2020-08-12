package main

import (
	"fmt"
	"net/http"
)

func headers(writer http.ResponseWriter, request *http.Request) {
	header := request.Header.Get("Accept-Encoding")
	fmt.Fprintln(writer, header)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}