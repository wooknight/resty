package main

import (
	"fmt"
	"kv/data"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloMuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Gorilla Mux\n"))
}

func friskMuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Frisky from Gorilla Mux\n"))
}
func businessMuxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	val, _ := data.Get(key)
	w.Write([]byte("Gorilla muxh business" + val + "\n"))
}

func ratingListMuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pictures \n Videos \n Text \n"))
}
func ratingMuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Business from Gorilla Mux\n"))
}

func main() {
	fmt.Println("Here we go")
	data.Put("121", "pav")
	r := mux.NewRouter()
	r.HandleFunc("/", helloMuxHandler)
	r.HandleFunc("/business/{key}", businessMuxHandler)
	r.HandleFunc("/user/rating/{type}", ratingListMuxHandler)
	r.HandleFunc("/user/rating/{id:[0-9]+}", ratingMuxHandler)
	log.Fatal(http.ListenAndServe(":8081", r))
}
