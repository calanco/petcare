package pet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Marshal JSON dat in Pet struct
func getPet(p *Pet, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ReadAll err: %v", err)
		log.Println(err)
		return
	}
	if err := json.Unmarshal(body, &p); err != nil {
		fmt.Fprintf(w, "Unmarshal err: %v", err)
		log.Println(err)
		return
	}
	if p.Nome == "" {
		fmt.Fprintf(w, "Nessun nome fornito")
		return
	}
}

// Store Pet in data.txt file
func storePet(p Pet, w http.ResponseWriter) {
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
