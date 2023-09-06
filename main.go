package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	HTTP_PORT := os.Getenv("HTTP_PORT")

	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr)
		fmt.Fprintf(w, "%q", html.EscapeString(r.RemoteAddr))
	})

	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(r.RemoteAddr)
		fmt.Fprintf(w, "Hit API!")
	})

	log.Fatal(http.ListenAndServe(":"+HTTP_PORT, nil))
}
