package food

import (
	"testing"

	"github.com/calanco/petcare/api/v1/pet"
)

// TestStruct contains the Pet to store and the expected wanted food
type TestStruct struct {
	Pet  pet.Pet
	Food string
}

// Tests defines the tests to run
var Tests = []TestStruct{
	{
		Pet: pet.Pet{
			Name:    "dog",
			Species: "dog",
		},
		Food: "dog food",
	},
	{
		Pet: pet.Pet{
			Name:    "cat",
			Species: "cat",
		},
		Food: "cat food",
	},
	{
		Pet: pet.Pet{
			Name:    "unknown",
			Species: "unknown",
		},
		Food: "Unknown species: no food",
	},
}

func TestGetFood(t *testing.T) {
	for _, test := range Tests {
		if getFood(test.Pet) != test.Food {
			t.Errorf("Test with %s doesn't pass", test.Pet.Name)
			continue
		}
		t.Logf("Test with %s passes", test.Pet.Name)
	}
}
