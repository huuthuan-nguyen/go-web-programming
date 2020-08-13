package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func writeExample(writer http.ResponseWriter, request *http.Request) {
	str := `<html>
	<head>
		<title>Go Web Programming</title>
	</head>
	<body>
		<h1>Hello World!</h1>
	</body>
	</html>`
	writer.Write([]byte(str))
}

func writeHeaderExample(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(501)
	fmt.Fprintln(writer, "No such service, try next door!")
}

func headerExample(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Location", "https://google.com")
	writer.WriteHeader(302)
}

func jsonExample(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	slcD := []string{"apple", "peach", "pear"}
	json, _ := json.Marshal(slcD)

	writer.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}