package pet

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// ListHandler serves HTTP requests at /api/pet/list endpoint
// It sends all saved pets to caller
func ListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json, err := json.Marshal(PetMap)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 500)
		logrus.Error(err)
		return
	}
	fmt.Fprintln(w, string(json))
	logrus.Info("Pets listed")
}
