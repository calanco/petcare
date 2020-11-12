package pet

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// TestDeletePet checks if created pets can be deleted
func TestDeletePet(t *testing.T) {
	testPet := Pet{
		Name:    "Bruce0",
		Species: "Dog",
	}

	CreatePet(&testPet)
	_, err := GetPet(string(testPet.Name))
	if err != nil {
		t.Errorf("%s doesn't pass", testPet.Name)
		return
	}
	err = DeletePet(testPet.Name)
	if err != nil {
		t.Errorf("%s doesn't pass", testPet.Name)
	}
	logrus.WithFields(logrus.Fields{
		"pet": testPet.Name,
	}).Info("Pet deleted")
}
