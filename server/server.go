package server

import (
	"github.com/gorilla/mux"
)

// Server creates a Gorilla Mux Router for the PetCare Server
func Server() *mux.Router {
	mux := mux.NewRouter()

	endpoints := GetEndpoints()

	for _, endpoint := range endpoints {
		mux.HandleFunc(endpoint.Path, endpoint.Function).Methods(endpoint.Method)
	}

	return mux
}
