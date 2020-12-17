package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/puxin71/talk-server/pkg"
)

// Loading the dataset from local json files
type FileLoader struct {
	Filename string
}

// Instantiate the file loader
func NewFileLoader(config pkg.ConfigProvider) FileLoader {
	dir := config.GetResourcePath()
	return FileLoader{Filename: filepath.Join(dir, "dataset.json")}
}

// Load all the dataset into memory
func (l FileLoader) GetAllTalks() ([]Talk, error) {
	var err error
	talks := make([]Talk, 0)

	// Open the json file
	jsonfile, err := os.Open(l.Filename)
	if err != nil {
		log.Println("failed to open file: ", l.Filename, ", error: ", err)
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

func GetJSONFile() string {
	return "dataset.json"
}
