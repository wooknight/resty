package server

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("IP ADDRESS %s ", GetIP(r))
		fmt.Fprintf(w, "IP ADDRESS %s", GetIP(r))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetIP(r *http.Request) string {
	var forwarded string
	if forwarded = r.Header.Get("X-FORWARDED-FOR"); forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
