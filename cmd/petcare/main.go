package main

import (
	"log"
	"net/http"

	"github.com/calanco/petcare/api/food"
	"github.com/calanco/petcare/api/home"
	"github.com/calanco/petcare/api/pet"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", home.Handler).Methods("GET")
	mux.HandleFunc("/api/pet/list", pet.ListHandler).Methods("GET")
	mux.HandleFunc("/api/pet/{name}", pet.GetHandler).Methods("GET")
	mux.HandleFunc("/api/pet", pet.CreateHandler).Methods("POST")
	mux.HandleFunc("/api/pet/{name}", pet.DeleteHandler).Methods("DELETE")
	mux.HandleFunc("/api/food/{name}", food.GetHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":80", mux))
}
