package pet

import "net/http"

// PutHandler serves PUT HTTP requests at /api/pet/{name} endpoint
func PutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	// TODO
}
