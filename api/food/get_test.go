package food

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/calanco/petcare/api/pet"
)

func TestGetHandler(t *testing.T) {
	// Storing test pets in PetMap
	for _, test := range Tests {
		if test.Pet.Name == "unknown" {
			continue
		}
		pet.PetMap[string(test.Pet.Name)] = test.Pet
	}

	fmt.Println(pet.PetMap)

	for _, test := range Tests {
		req, err := http.NewRequest("GET", "/api/food/{name}", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetHandler)
		q := req.URL.Query()
		q.Add("name", string(test.Pet.Name))
		req.URL.RawQuery = q.Encode()
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("%s: handler returned wrong status code: got %v want %v", test.Pet.Name, status, http.StatusOK)
		}

		if rr.Body.String() != test.Food {
			t.Errorf("%s: handler returned unexpected body: got %v want %v", test.Pet.Name, rr.Body.String(), test.Food)
		}
	}
}
