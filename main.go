package main

import (
	"log"
	"net/http"

	"github.com/calanco/petcare/pkg/home"
	"github.com/calanco/petcare/pkg/pet"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", home.HomeHandler).Methods("GET")
	mux.HandleFunc("/api/pet/store", pet.StoreHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", mux))
}
