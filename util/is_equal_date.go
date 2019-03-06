package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func IsEqualDate(offDate []string) bool {
	splitDate := strings.Split(offDate[0], "-")
	year := splitDate[0]
	month := splitDate[1]

	p, _ := os.Getwd()
	inputJSON := filepath.Join(p, "www", "api", "v1", year, month+".json")

	jsonString, err := ioutil.ReadFile(inputJSON)
	if err != nil {
		log.Print(err)
		return false
	}

	jsonData := new(Api)
	if err = json.Unmarshal(jsonString, jsonData); err != nil {
		log.Fatal(err)
	}

	return reflect.DeepEqual(offDate, jsonData.Date)
}
