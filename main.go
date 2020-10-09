package main

import (
	"log"
	"net/http"

	"github.com/calanco/petcare/pkg/home"
	"github.com/calanco/petcare/pkg/pet"
)

func main() {
	homeHandler := &home.Home{}
	petHandler := &pet.Pet{}

	mux := http.NewServeMux()
	mux.Handle("/", homeHandler)
	mux.Handle("/pet", petHandler)
	log.Fatal(http.ListenAndServe(":80", mux))
}
