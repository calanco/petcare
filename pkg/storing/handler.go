package storing

import (
	"fmt"
	"log"
	"net/http"
)

// Function to serve HTTP requets at /api/store endpoint
func StoreHandler(w http.ResponseWriter, r *http.Request) {
	p := &Pet{}
	if err := p.parseJSON(w, r); err != nil {
		fmt.Fprintf(w, "Err: %v", err)
		log.Println(err)
		return
	}
	if p.Nome != "" {
		p.storePet(w)
		fmt.Fprintf(w, "Salvato: %v", p)
		return
	}
	fmt.Fprintf(w, "Non salvato: %v", p)
}
