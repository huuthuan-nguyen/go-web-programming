package main

import (
	"net/http"
	"fmt"
)

func process(writer http.ResponseWriter, request *http.Request) {
	// request.ParseForm()
	// request.ParseMultipartForm(1024)
	// fmt.Fprintln(writer, request.FormValue("first_name"))
	fmt.Fprintln(writer, request.PostFormValue("first_name"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}