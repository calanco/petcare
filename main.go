package main

import (
	"log"
	"net/http"

	"github.com/calanco/petcare/pkg/home"
	"github.com/calanco/petcare/pkg/storing"
)

func main() {
	homeHandler := &home.Home{}
	storeHandler := &storing.Pet{}

	mux := http.NewServeMux()
	mux.Handle("/", homeHandler)
	mux.Handle("/store", storeHandler)
	log.Fatal(http.ListenAndServe(":80", mux))
}
