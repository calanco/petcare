package pet

import (
	"bufio"
	"encoding/json"
	"os"
)

// DATAPATH stores the text file path where to store pets possibly
const DATAPATH = "data.txt"

// createFileIfNotExists creates data.txt files if it doesn't exists
func createFileIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		defer f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// writeOnFile writes Pet on data.txt file
func writeOnFile(p Pet, path string) error {
	f, err := os.OpenFile(path, os.O_APPEND, 0644)
	defer f.Close()

	out, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = f.WriteString(string(out) + "\n")
	if err != nil {
		return err
	}

	return nil
}

// readFromFile reads from data.txt file, line by line
func readFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return content, nil
}

// CleanPets destroys all the stored pets
func CleanPets() {
	PetMap = make(map[string]Pet)
}
