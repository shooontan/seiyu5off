package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/shooontan/seiyu5off/util"
)

func main() {
	apiData := util.Scrape()

	// months
	for month := range apiData.Date {

		splitDate := strings.Split(apiData.Date[month][0], "-")
		yearDate := splitDate[0]
		monthDate := splitDate[1]

		jsonData := util.ReadJSON(yearDate, monthDate)

		if isEqual := util.IsEqualDate(apiData.Date[month], jsonData); !isEqual {
			util.Generate(util.Api{
				Updated: apiData.Updated,
				Date:    util.MergeDate(apiData.Date[month], jsonData),
			}, yearDate, monthDate+".json")
		}
	}

	// years
	p, _ := os.Getwd()
	outputDir := filepath.Join(p, "www", "api", "v1", strings.Repeat("[0-9]", 4))
	files, err := filepath.Glob(outputDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		base := filepath.Base(file)
		yearJSONDate := util.ReadYearJSON(base)
		util.Generate(util.Api{
			Updated: yearJSONDate.Updated,
			Date:    yearJSONDate.Date,
		}, "", base+".json")
	}
}
