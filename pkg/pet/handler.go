package pet

import (
	"net/http"
)

// Function to serve HTTP requets at /pet endpoint
func (p Pet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		getPet(&p, w, r)
		if p.Nome != "" {
			storePet(p, w)
		}
	}
}
