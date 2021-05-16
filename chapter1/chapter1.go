package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"html/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func templater() {

	http.ListenAndServe(":80", nil)
}

type UptimeHandler struct {
	Started time.Time
}

func NewUptimeHandler() UptimeHandler {
	return UptimeHandler{Started: time.Now()}
}

func (h UptimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current uptime : %s", time.Since(h.Started))
}

type SecretTokenHandler struct {
	next   http.Handler
	secret string
}

func (s SecretTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("secret") == s.secret {
		s.next.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "I am getting out")
	}
	pucchiHandler := func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "http://www.smcl.org", http.StatusPermanentRedirect)
	}
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/kate", pucchiHandler)
	http.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://www.google.com", http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello from /users")
		http.Error(w, "fucked", http.StatusNotFound)
	})
	http.Handle("/bhosada", SecretTokenHandler{
		next:   NewUptimeHandler(),
		secret: "mysecret",
	})

	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
