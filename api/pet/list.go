package pet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ListHandler serves HTTP requests at /api/pet/list endpoint
// It sends all saved pets to caller
func ListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json, err := json.Marshal(PetMap)
	if err != nil {
		http.Error(w, error.Error(err), 500)
		return
	}
	fmt.Fprintln(w, string(json))
	log.Println("List requested:", string(json))
}
