package main

import (
	"net/http"
	"math/rand"
	"html/template"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func processRandNumber(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("template.html")
	rand.Seed(time.Now().Unix())
	t.Execute(writer, rand.Intn(10) > 5)
}

func processDaysOfWeek(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("days_of_week.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(writer, daysOfWeek)
}

func processSetAction(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("set_action.html")
	t.Execute(writer, "hello")
}

func processIncludeAction(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("master.html", "sub.html")
	t.Execute(writer, "Hello World!")
}

func processCustomFunction(writer http.ResponseWriter, request *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("custom_function.html").Funcs(funcMap)
	t, _ = t.ParseFiles("custom_function.html")
	t.Execute(writer, time.Now())
}

func processContextAwareness(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("context_awareness.html")
	context := `I asked: <i>What's up?</i>`
	t.Execute(writer, template.HTML(context))
}

func processNestingTemplate(writer http.ResponseWriter, request *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
	} else {
		t, _ = template.ParseFiles("layout.html")
	}
	t.ExecuteTemplate(writer, "layout", nil)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", processRandNumber)
	http.HandleFunc("/days_of_week", processDaysOfWeek)
	http.HandleFunc("/set_action", processSetAction)
	http.HandleFunc("/include_action", processIncludeAction)
	http.HandleFunc("/custom_function", processCustomFunction)
	http.HandleFunc("/context_awareness", processContextAwareness)
	http.HandleFunc("/nesting_template", processNestingTemplate)
	server.ListenAndServe()
}