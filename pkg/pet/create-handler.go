package pet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Function to serve HTTP requets at /api/pet/create endpoint
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	p := Pet{}
	if err := p.parseJSON(w, r); err != nil {
		fmt.Fprintf(w, "Err: %v", err)
		log.Println(err)
		return
	}
	PetMap[p.Name] = p
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
	if p.Name == "" {
		return errors.New("No name defined")
	}
	if p.Species == "" {
		return errors.New("No species defined")
	}
	return nil
}
