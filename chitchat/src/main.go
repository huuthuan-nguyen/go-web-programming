package main

import (
	"data"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	} else {
		error_message(writer, request, "Cannot get threads")
	}
}

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)

	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func generateHTML(writer http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

func main() {
	mux := http.NewServeMux() // this is like a controller in PHP
	files := http.FileServer(http.Dir("/public")) // create handler function, like an action of controller.
	
	mux.Handle("/static/", http.StripPrefix("/static/", files)) // put the handler function to handle, like register an action to a route.

	mux.HandleFunc("/", index) // default endpoint. callback as parameter.
	mux.HandleFunc("/err", err)
	// mux.handleFunc("/login", login)
	// mux.HandleFunc("/logout", logout)
	// mux.HandleFunc("/signup", signup)
	// mux.HandleFunc("/signup_account", signupAccount)
	// mux.HandleFunc("/authenticate", authenticate)

	// mux.HandleFunc("/thread/new", newThread)
	// mux.HandleFunc("/thread/create", createThread)
	// mux.HandleFunc("/thread/post", postThread)
	// mux.HandleFunc("/thread/read", readThread)


	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

type Server struct {
	Addr string
	// Handler mux
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	MaxHeaderBytes int
	// TLSConfig *tls.Config
	// TLSNextProto map[string]func(*Server, *tls.Conn, Hanlder)
	// ConnState func(net.Conn, ConnState)
	//ErrorLog *log.Logger
}