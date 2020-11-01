package server

import (
	"net/http"

	"github.com/calanco/petcare/api/food"
	"github.com/calanco/petcare/api/home"
	"github.com/calanco/petcare/api/pet"
)

// Endpoint defines the attributes of a PetCare endpoint
type Endpoint struct {
	Path     string
	Function func(http.ResponseWriter, *http.Request)
	Method   string
}

// GetEndpoints returns all the needed PetCare endpoints
func GetEndpoints() []Endpoint {
	endpoints := []Endpoint{
		{
			Path:     "/",
			Function: home.Handler,
			Method:   "GET",
		},
		{
			Path:     "/api/pet/list",
			Function: pet.ListHandler,
			Method:   "GET",
		},
		{
			Path:     "/api/pet/{name}",
			Function: pet.GetHandler,
			Method:   "GET",
		},
		{
			Path:     "/api/pet",
			Function: pet.CreateHandler,
			Method:   "POST",
		},
		{
			Path:     "/api/pet/{name}",
			Function: pet.DeleteHandler,
			Method:   "DELETE",
		},
		{
			Path:     "/api/food/{name}",
			Function: food.GetHandler,
			Method:   "GET",
		},
	}
	return endpoints
}
