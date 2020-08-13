package main

import (
	"net/http"
	"fmt"
)

func setCookie(writer http.ResponseWriter, request *http.Request) {
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name: "second_cookie",
		Value: "Manning Publication Co",
		HttpOnly: true,
	}

	// writer.Header().Set("Set-Cookie", c1.String())
	// writer.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(writer, &c1)
	http.SetCookie(writer, &c2)
}

func getCookie(writer http.ResponseWriter, request *http.Request) {
	// regular method
	// header := request.Header["Cookie"]
	// fmt.Fprintln(writer, header)
	// advanced method
	c1, err := request.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(writer, "Can not get the first cookie")
	}
	cs := request.Cookies()
	fmt.Fprintln(writer, c1)
	fmt.Fprintln(writer, cs)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}