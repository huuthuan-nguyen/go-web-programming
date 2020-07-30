package main

import (
	"data"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var _, filename, _, _ = runtime.Caller(0)
var rootPath = filepath.Dir(filename)

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
		errorMessage(writer, request, "Cannot get threads")
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

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	for _, file := range filenames {
		templatePath := rootPath + string(os.PathSeparator) + "templates" + string(os.PathSeparator) + "%s.html"
		files = append(files, fmt.Sprintf(templatePath, file))
	}

	templates := template.Must(template.ParseFiles(files...))
	_ = templates.ExecuteTemplate(writer, "layout", data)
}

func main() {
	// handle the static assets.
	mux := http.NewServeMux() // this is like a controller in PHP
	files := http.FileServer(http.Dir(rootPath + string(os.PathSeparator) + "public")) // create handler function, like an action of controller.

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
		ReadTimeout: time.Duration(10 * int64(time.Second)),
		WriteTimeout: time.Duration(600 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe()
}

type Server struct {
	Addr string
	// Handler mux
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	MaxHeaderBytes int
	// TLSConfig *tls.Config
	// TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
	// ConnState func(net.Conn, ConnState)
	//ErrorLog *log.Logger
}