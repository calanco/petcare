package pet

import (
	"testing"
)

// TestGetPet checks if all pets that can be created are able to be retrieved
func TestGetPet(t *testing.T) {
	CleanPets()

	for _, testPet := range TestPets {
		CreatePet(&testPet.pet)
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
