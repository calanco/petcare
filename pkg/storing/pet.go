package storing

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Pet defines pet attributes
type Pet struct {
	Nome  string `json:"nome"`
	Breed Breed  `json:"breed"`
	Size  Size   `json:"size"`
}

// Marshal JSON dat in Pet struct
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
	createFileIfNotExists(path)
	f, err := os.OpenFile(path, os.O_APPEND, 0644)
	defer f.Close()

	out, err := json.Marshal(p)
	if err != nil {
		fmt.Fprintf(w, "Marshal err: %v", err)
		log.Println(err)
	}

	_, err = f.WriteString(string(out) + "\n")
	if err != nil {
		fmt.Fprintf(w, "WriteString err: %v", err)
		log.Println(err)
	}
}
