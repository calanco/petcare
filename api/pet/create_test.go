package pet

import (
	"testing"
)

// TestCreatePet checks that pet creations work properly
func TestCreatePet(t *testing.T) {
	CleanPets()

	for _, testPet := range TestPets {
		err := CreatePet(&testPet.pet)
		if err != nil && testPet.ok {
			t.Errorf("%s doesn't pass", testPet.testName)
			continue
		}
		if err == nil && !testPet.ok {
			t.Errorf("%s doesn't pass", testPet.testName)
			continue
		}
		t.Logf("Test with %s passes", testPet.testName)
	}

}
