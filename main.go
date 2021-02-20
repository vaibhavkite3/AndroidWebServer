package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func main() {
	http.Handle("/", http.FileServer(rice.MustFindBox("dist").HTTPBox()))
	http.ListenAndServe(":8080", nil)
}
