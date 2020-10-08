package main

import (
	"log"
	"net/http"

	"github.com/calanco/petcare/pkg/handler"
)

func main() {
	homeHandler := &handler.Home{}

	mux := http.NewServeMux()
	mux.Handle("/", homeHandler)
	log.Fatal(http.ListenAndServe(":80", mux))
}
