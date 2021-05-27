package main

import (
	"fmt"
	"io"
	"kv/data"
	"log"
	"net/http"
	"os"

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
func businessCreateMuxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = data.Put(key, (string)(value))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ratingListMuxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%v\n", vars)
	key := vars["type"]
	val, _ := data.Get(key)
	w.Write([]byte("Pictures \n Videos \n Text \n" + val))
}
func ratingMuxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	val, _ := data.Get(key)
	w.Write([]byte("Hello Business from Gorilla Mux\n" + val))
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	maxMem := (int64)(16 << 24)
	err := r.ParseMultipartForm(maxMem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("WTF - why am I missing data\n" + err.Error()))
		return
	}
	data := r.MultipartForm
	files := data.File["files"]
	fmt.Fprintf(w, "%v", files)

	for _, fh := range files {
		f, err := fh.Open()
		defer f.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Could not open uploaded file\n" + err.Error()))
			return
		}
		out, err := os.Create("C:\\Users\\ram_n\\Downloads\\" + fh.Filename)
		defer out.Close()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Could not create temp file\n" + err.Error()))
			return
		}
		_, err = io.Copy(out, f)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Could not save temp file\n" + err.Error()))
			return
		}
		fmt.Fprintln(w, "Upload completed")
	}

}

func main() {
	fmt.Println("Here we go")
	data.Put("121", "pav")
	data.Put("video", "pucchi")
	r := mux.NewRouter()
	r.HandleFunc("/", helloMuxHandler)
	r.HandleFunc("/business/{key}", businessMuxHandler).Methods("GET")
	r.HandleFunc("/business/{key}", businessCreateMuxHandler).Methods("PUT")
	r.HandleFunc("/user/image", imageHandler).Methods("POST")
	r.HandleFunc("/user/rating/{type}", ratingListMuxHandler)
	r.HandleFunc("/user/rating/{id:[0-9]+}", ratingMuxHandler)
	log.Fatal(http.ListenAndServe(":8081", r))
}
