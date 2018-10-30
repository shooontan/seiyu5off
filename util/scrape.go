package util

import (
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Scrape() Scraping {
	// Webページ
	targetURL := "https://www.seiyu.co.jp/service/5off/"

	// リクエスト
	res, err := http.Get(targetURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// HTML読み込み
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	// 5%offの月日を取得
	offDateMap := DateMap{}
	doc.Find(".off_calendar_item").Each(func(i int, s *goquery.Selection) {
		offMonth := s.Find(".off_calendar_month").Text()
		offMonth = strings.Trim(offMonth, "/")

		offDay := s.Find(".off_calendar_day").Text()

		offTimeYear := Str2int(GetYear(offMonth))
		offTimeMonth := time.Month(Str2int(offMonth))
		offTimeDay := Str2int(offDay)

		offDate := time.Date(offTimeYear, offTimeMonth, offTimeDay, 0, 0, 0, 0, time.Local).Format("2006-01-02")
		offDateMap[offMonth] = append(offDateMap[offMonth], offDate)
	})

	// 日付をソートする
	for month := range offDateMap {
		sort.SliceStable(offDateMap[month], func(i, j int) bool {
			return offDateMap[month][i] < offDateMap[month][j]
		})
	}

	return Scraping{
		Updated: time.Now().Format("2006-01-02 15:04:05"),
		Date:    offDateMap,
	}
}
