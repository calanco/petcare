package pet

import (
	"testing"
)

type TestPet struct {
	pet Pet
	ok  bool
}

func TestGetPet(t *testing.T) {
	testPets := []TestPet{
		{
			Pet{
				Name:    "Bruce",
				Species: "Dog",
			},
			true,
		},
		{
			Pet{
				Name:    "Tina",
				Species: "Dog",
			},
			true,
		},
		{
			Pet{
				Name:    "Kira",
				Species: "Cat",
			},
			true,
		},
	}

	for _, testPet := range testPets {
		CreatePet(testPet.pet)
		_, err := GetPet(string(testPet.pet.Name))
		if err != nil {
			if testPet.ok {
				t.Error(testPet.pet.Name)
				continue
			}
		} else {
			if !testPet.ok {
				t.Error(testPet.pet.Name)
				continue
			}
		}
		t.Logf("Test with %s passes", testPet.pet.Name)
	}
}
