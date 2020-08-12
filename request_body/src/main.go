package main

import (
	"fmt"
	"net/http"
)

func body(writer http.ResponseWriter, request *http.Request) {
	len := request.ContentLength
	body := make([]byte, len)
	request.Body.Read(body)
	fmt.Fprintln(writer, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/body", body)
	server.ListenAndServe()
}