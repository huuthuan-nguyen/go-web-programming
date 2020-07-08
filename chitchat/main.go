package main

import (
	"net/http"
)

/* index func */
// func index(w http.ResponseWriter, r *http.Request) {
// 	threads, err := data.Threads(); if err == nil {
// 		_, err := session(w, r)
// 		public_tmpl_files := []string{"templates/layout.html",
// 										"templates/public.navbar.html",
// 										"templates/index.html"}
// 		private_tmpl_files := []string{"templates/layout.html",
// 										"templates/private.navbar.html",
// 										"templates/index.html"}
// 		var templates *template.Template
// 		if err != nil {
// 			templates = template.Must(template.ParseFiles(private_tmpl_files...))
// 		} else {
// 			templates = template.Must(template.ParseFiles(public_tmpl_files...))
// 		}
// 		templates.ExecuteTemplate(w, "layout", threads)
// 	}
// }
func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}

func generateHTML(w http.ResponseWrite, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(write, "layout", data)
}


/* main func */
func main() {
	mux := http.NewServeMux() // this is like a controller in PHP
	files := http.FileServer(http.Dir("/public")) // create handler function, like an action of controller.
	
	mux.Handle("/static/", http.StripPrefix("/static/", files)) // put the handler function to handle, like register an action to a route.

	mux.HandleFunc("/", index) // default endpoint. callback as parameter.
	mux.HandleFunc("/err", err)
	mux.handleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)


	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handle: mux,
	}
	server.ListenAndServe()
}