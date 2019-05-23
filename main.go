package main

import (
	"strings"

	"github.com/shooontan/seiyu5off/util"
)

func main() {
	apiData := util.Scrape()

	for month := range apiData.Date {

		splitDate := strings.Split(apiData.Date[month][0], "-")
		jsonData := util.ReadJSON(splitDate[0], splitDate[1])

		if isEqual := util.IsEqualDate(apiData.Date[month], jsonData); !isEqual {
			util.Generate(util.Api{
				Updated: apiData.Updated,
				Date:    util.MergeDate(apiData.Date[month], jsonData),
			})
		}
	}
}
