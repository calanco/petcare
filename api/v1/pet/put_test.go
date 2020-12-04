package pet

import (
	"testing"
)

// TestPutPet checks that pet updates work properly
func TestPutPet(t *testing.T) {
	CleanPets()

	for _, testPet := range TestPets {
		err := PutPet(&testPet.pet)
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
