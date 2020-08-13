package main

import (
	"net/http"
	"html/template"
)

func process(writer http.ResponseWriter, request *http.Request) {
	// t, _ := template.ParseFiles("template.html")
	// parse with ParseGlob method
	// t, _ := template.ParseGlob("*.html")
	// t := template.Must(template.ParseFiles("template.html"))
	t := template.Must(template.ParseFiles("template.html", "other.html"))
	// t.Execute(writer, "Hello World!")
	t.ExecuteTemplate(writer, "other.html", "Hello World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}