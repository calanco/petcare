package pet

import (
	"testing"
)

// TestPet defines the struct to run test with
type TestPet struct {
	testName string // defines the name of the test
	pet      Pet    // represents the pet to test
	ok       bool   // defines if the pet has to be correctly created
}

// TestGetPet checks if all pets that can be created are able to be retrieved
func TestGetPet(t *testing.T) {
	testPets := []TestPet{
		{
			"Working test",
			Pet{
				Name:    "Bruce0",
				Species: "Dog",
			},
			true,
		},
		{
			"Wrong breed test",
			Pet{
				Name:    "Tina0",
				Species: "Dog",
				Breed:   "Unknown",
			},
			false,
		},
		{
			"Working date test",
			Pet{
				Name:    "Tina1",
				Species: "Dog",
				Date:    "2006-May-30",
			},
			true,
		},
		{
			"Wrong date test",
			Pet{
				Name:    "Tina2",
				Species: "Dog",
				Date:    "Wrong date",
			},
			false,
		},
		{
			"Working size",
			Pet{
				Name:    "Kira0",
				Species: "Cat",
				Size:    "Medium",
			},
			true,
		},
		{
			"Working weight",
			Pet{
				Name:    "Kira1",
				Species: "Cat",
				Weight:  2.5,
			},
			true,
		},
		{
			"Wrong species",
			Pet{
				Name:    "Unknown0",
				Species: "Unknown",
			},
			false,
		},
	}

	for _, testPet := range testPets {
		CreatePet(testPet.pet)
		_, err := GetPet(string(testPet.pet.Name))
		if err != nil {
			if testPet.ok {
				t.Errorf("%s doesn't pass", testPet.testName)
				continue
			}
		} else {
			if !testPet.ok {
				t.Errorf("%s doesn't pass", testPet.testName)
				continue
			}
		}
		t.Logf("Test with %s passes", testPet.testName)
	}
}
