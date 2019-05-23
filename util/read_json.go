package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ReadJSON reads JSON files.
func ReadJSON(year string, month string) *Api {
	jsonData := new(Api)

	p, _ := os.Getwd()
	inputJSON := filepath.Join(p, "www", "api", "v1", year, month+".json")

	jsonString, err := ioutil.ReadFile(inputJSON)
	if err != nil {
		log.Print(err)
		return jsonData
	}

	if err = json.Unmarshal(jsonString, jsonData); err != nil {
		log.Print(err)
	}

	return jsonData
}
