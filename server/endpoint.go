package server

import (
	"net/http"

	"github.com/calanco/petcare/api/v1/food"
	"github.com/calanco/petcare/api/v1/home"
	"github.com/calanco/petcare/api/v1/pet"
	"github.com/calanco/petcare/api/v1/pets"
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
		// Home endpoints
		{
			Path:     "/",
			Function: home.Handler,
			Method:   "GET",
		},
		// Pet endpoints
		{
			Path:     "/api/pet/{name}",
			Function: pet.GetHandler,
			Method:   "GET",
		},
		{
			Path:     "/api/pet",
			Function: pet.PostHandler,
			Method:   "POST",
		},
		{
			Path:     "/api/pet/{name}",
			Function: pet.PutHandler,
			Method:   "PUT",
		},
		{
			Path:     "/api/pet/{name}",
			Function: pet.DeleteHandler,
			Method:   "DELETE",
		},
		// Pets endpoints
		{
			Path:     "/api/pets",
			Function: pets.GetHandler,
			Method:   "GET",
		},
		// Food endpoints
		{
			Path:     "/api/food/{name}",
			Function: food.GetHandler,
			Method:   "GET",
		},
	}
	return endpoints
}
