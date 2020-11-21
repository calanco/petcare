package food

import (
	"github.com/calanco/petcare/api/v1/pet"
)

// Food is a string defining the food for a determined pet
type Food string

// Get food for p pet
func getFood(p pet.Pet) string {
	switch p.Species {
	case "dog":
		return "dog food"
	case "cat":
		return "cat food"
	default:
		return "Unknown species: no food"
	}
}
