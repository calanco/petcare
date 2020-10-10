package pet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Function to serve HTTP requets at /api/pet/store endpoint
func StoreHandler(w http.ResponseWriter, r *http.Request) {
	p := &Pet{}
	if err := p.parseJSON(w, r); err != nil {
		fmt.Fprintf(w, "Err: %v", err)
		log.Println(err)
		return
	}
	if p.Nome != "" {
		p.storePet(w)
		fmt.Fprintln(w, p)
		return
	}
	fmt.Fprintf(w, "Non salvato: %v", p)
}

// Marshal JSON data in Pet struct
func (p *Pet) parseJSON(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &p); err != nil {
		return err
	}
	if p.Nome == "" {
		return errors.New("Nessun nome fornito")
	}
	return nil
}

// Store Pet in data.txt file
func (p *Pet) storePet(w http.ResponseWriter) {
	path := "data.txt"
	if err := createFileIfNotExists(path); err != nil {
		fmt.Fprintf(w, "Err: %v", err)
		log.Println(err)
	}
	if err := writeOnFile(*p, path); err != nil {
		fmt.Fprintf(w, "Err: %v", err)
		log.Println(err)
	}

}
