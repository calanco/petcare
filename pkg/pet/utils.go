package pet

import (
	"log"
	"os"
)

// Func to create data.txt files if it doesn't exists
func createFileIfNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
