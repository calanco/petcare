package pet

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Func to create data.txt files if it doesn't exists
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

// Func to write Pet on data.txt file
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

// Func to read from data.txt file, line by line
func readFromFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
