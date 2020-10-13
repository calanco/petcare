package food

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	pet "github.com/calanco/petcare/pkg/pet"
)

type petName struct {
	Name string `json:name`
}

// Function to serve HTTP requets at /api/food/get endpoint
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/raw")

	name, err := getPetName(r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	p, err := pet.GetPet(name)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	food := getFood(p)
	fmt.Fprintln(w, food)
}

// Get pet name from request
func getPetName(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	pn := petName{}
	if err := json.Unmarshal(body, &pn); err != nil {
		return "", err
	}

	if pn.Name == "" {
		return "", errors.New("No name set")
	}

	return pn.Name, nil
}
