package pet

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Function to serve HTTP requets at /api/pet/list endpoint
// It sends all saved pets to caller
func ListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json, err := json.Marshal(PetMap)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, string(json))
	log.Println("Listed:", string(json))
}
