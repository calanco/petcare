package pets

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/calanco/petcare/api/pet"
	"github.com/sirupsen/logrus"
)

// GetHandler serves HTTP requests at /api/pets endpoint
// It lists all saved pets to caller
func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json, err := json.Marshal(pet.PetMap)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		logrus.Error(err)
		return
	}
	fmt.Fprintln(w, string(json))
	logrus.Info("Pets listed")
}
