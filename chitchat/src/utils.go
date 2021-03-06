package main

import (
	"data"
	"errors"
	"log"
	"net/http"
	"strings"
)

var logger *log.Logger

// check if the user is logged in and has a session, if not err is not nil
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("invalid session")
		}
	}
	return 
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

// convenience function to redirect to the error message page
func errorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}