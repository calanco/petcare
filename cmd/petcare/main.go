package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/calanco/petcare/api/food"
	"github.com/calanco/petcare/api/home"
	"github.com/calanco/petcare/api/pet"
	"github.com/gorilla/mux"
)

var port string

func init() {
	flag.StringVar(&port, "port", "80", "Petcare listening port")
	flag.Parse()
}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", home.Handler).Methods("GET")
	mux.HandleFunc("/api/pet/list", pet.ListHandler).Methods("GET")
	mux.HandleFunc("/api/pet/{name}", pet.GetHandler).Methods("GET")
	mux.HandleFunc("/api/pet", pet.CreateHandler).Methods("POST")
	mux.HandleFunc("/api/pet/{name}", pet.DeleteHandler).Methods("DELETE")
	mux.HandleFunc("/api/food/{name}", food.GetHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
