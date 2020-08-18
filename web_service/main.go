package main

import (
	"net/http"
	"path"
	"encoding/json"
	"strconv"
)

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/posts/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(writer http.ResponseWriter, request *http.Request) {
	var err error
	switch request.Method{
	case "GET":
		err = handleGet(writer, request)
	case "POST":
		err = handlePost(writer, request)
	case "PUT":
		err = handlePut(writer, request)
	case "DELETE":
		err = handleDelete(writer, request)
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(writer http.ResponseWriter, request *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(request.URL.Path))
	if err != nil {
		return
	}
	post := Post{
		Id: id,
		Title: "Hello World!",
		Content: "This is content",
	}
	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(output)
	return
}

func handlePost(writer http.ResponseWriter, request *http.Request) (err error) {
	len := request.ContentLength
	body := make([]byte, len)
	request.Body.Read(body)

	var post Post
	json.Unmarshal(body, &post)

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(201)
	writer.Write(output)
	return
}

func handlePut(writer http.ResponseWriter, request *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(request.URL.Path))
	if err != nil {
		return
	}

	post := Post{
		Id: id,
		Title: "Test",
		Content: "Test",
	}
	 
	len := request.ContentLength
	body := make([]byte, len)
	request.Body.Read(body)
	json.Unmarshal(body, &post)

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(output)
	return
}

func handleDelete(writer http.ResponseWriter, request *http.Request) (err error) {
	writer.WriteHeader(204)
	return
}