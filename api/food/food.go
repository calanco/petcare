package food

import (
	"fmt"

	"github.com/calanco/petcare/api/pet"
)

// Food is a string defining the food for a determined pet
type Food string

// Get food for p pet
func getFood(p pet.Pet) string {
	fmt.Println(p)
	switch p.Species {
	case "dog":
		return "dog food"
	case "cat":
		return "cat food"
	default:
		return "Unknown species: no food"
	}
}
