package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/ping"), ping)

	http.ListenAndServe("localhost:7771", mux)
}
