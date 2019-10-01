package main

import (
	"net/http"
	"os"
	"strings"
)

func root(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", root)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
