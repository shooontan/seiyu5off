package main

import (
	"github.com/shooontan/seiyu5off/util"
)

func main() {
	apiData := util.Scrape()

	for month := range apiData.Date {
		if isEqual := util.IsEqualDate(apiData.Date[month]); !isEqual {
			util.Generate(util.Api{
				Updated: apiData.Updated,
				Date:    apiData.Date[month],
			})
		}
	}
}
