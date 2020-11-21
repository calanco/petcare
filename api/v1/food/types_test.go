package food

import (
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
