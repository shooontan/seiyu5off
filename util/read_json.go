package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
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

func ReadYearJSON(year string) *Api {
	yearDate := new(Api)

	p, _ := os.Getwd()
	jsonPath := filepath.Join(p, "www", "api", "v1", year, "*.json")

	files, err := filepath.Glob(jsonPath)
	if err != nil {
		panic(err)
	}

	for _, filePath := range files {
		fileText, err := ioutil.ReadFile(filePath)

		if err != nil {
			log.Print(err)
			continue
		}

		jsonData := new(Api)
		if json.Unmarshal(fileText, jsonData) != nil {
			log.Print(err)
			continue
		}

		if len(yearDate.Updated) == 0 {
			yearDate.Updated = jsonData.Updated
		}

		yearUpdated, _ := time.Parse("2006-01-02 15:04:05", yearDate.Updated)
		jsonUpdated, _ := time.Parse("2006-01-02 15:04:05", jsonData.Updated)

		if yearUpdated.Before(jsonUpdated) {
			yearDate.Updated = jsonData.Updated
		}

		yearDate.Date = append(yearDate.Date, jsonData.Date...)
	}

	sort.SliceStable(yearDate.Date, func(i, j int) bool {
		return yearDate.Date[i] < yearDate.Date[j]
	})

	return yearDate
}
