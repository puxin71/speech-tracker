package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Loading the dataset from local json files
type FileLoader struct {
	filename string
}

// Instantiate the file loader
func NewFileLoader(file string) FileLoader {
	return FileLoader{filename: file}
}

// Load all the dataset into memory
func (l FileLoader) GetAll() ([]Talk, error) {
	var err error
	talks := make([]Talk, 0)

	// Open the json file
	jsonfile, err := os.Open(l.filename)
	if err != nil {
		log.Println("failed to open file: ", l.filename, ", error: ", err)
	}
	defer jsonfile.Close()

	// Read the data as a byte array and unmarshall the array
	// into our predefined data models
	byteValue, _ := ioutil.ReadAll(jsonfile)
	err = json.Unmarshal(byteValue, &talks)
	if err != nil {
		log.Println("json unmarshal error: ", err)
	}

	// Set the correct role for the speaker and attendees
	updateRole(talks)

	return talks, err
}

func updateRole(talks []Talk) {
	for i := range talks {
		talks[i].Speaker.Role = SPEAKER
	}
}
