package main

import (
	"net/http"
	"fmt"
	"time"
	"encoding/base64"
)

func setMessage(writer http.ResponseWriter, request *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}

	http.SetCookie(writer, &c)
}

func showMessage(writer http.ResponseWriter, request *http.Request) {
	c, err := request.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(writer, "No message found.")
		} else {
			rc := http.Cookie{
				Name: "flash",
				MaxAge: -1,
				Expires: time.Unix(1, 0),
			}
			http.SetCookie(writer, &rc)
			val, _ := base64.URLEncoding.DecodeString(c.Value)
			fmt.Fprintln(writer, string(val))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)

	server.ListenAndServe()
}