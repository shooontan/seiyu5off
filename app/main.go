package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// OffDateTime is Web API
type OffDateTime struct {
	Updated string   `json:"updated"`
	Date    []string `json:"date"`
}

func str2int(t string) int {
	i, _ := strconv.Atoi(t)
	return i
}

func getNowYear() string {
	return time.Now().Format("2006")
}

func scrapeOffDateTime() OffDateTime {
	nowYear := str2int(getNowYear())

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
	offDates := []string{}
	doc.Find(".off_calendar_item").Each(func(i int, s *goquery.Selection) {
		offMonth := s.Find(".off_calendar_month").Text()
		offMonth = strings.Trim(offMonth, "/")
		offDay := s.Find(".off_calendar_day").Text()
		offtime := time.Date(nowYear, time.Month(str2int(offMonth)), str2int(offDay), 0, 0, 0, 0, time.Local).Format("2006-01-02")
		offDates = append(offDates, offtime)
	})

	return OffDateTime{
		Updated: time.Now().Format("2006-01-02 15:04:05"),
		Date:    offDates,
	}
}

func isEqualDate(offDate []string) bool {
	nowMonth := time.Now().Format("01")
	p, _ := os.Getwd()
	inputJSON := filepath.Join(p, "dist", "api", getNowYear(), nowMonth+".json")

	jsonString, err := ioutil.ReadFile(inputJSON)
	if err != nil {
		log.Print(err)
		return false
	}

	jsonData := new(OffDateTime)
	if err = json.Unmarshal(jsonString, jsonData); err != nil {
		log.Fatal(err)
	}

	return reflect.DeepEqual(offDate, jsonData.Date)
}

func generateOffDate(apiData OffDateTime) {
	nowMonth := time.Now().Format("01")
	p, _ := os.Getwd()

	// json変換
	outputJSON, err := json.Marshal(apiData)
	if err != nil {
		log.Fatal("json encode error")
		panic(err)
	}
	out := new(bytes.Buffer)
	json.Indent(out, outputJSON, "", "    ")

	outputDir := filepath.Join(p, "dist", "api", getNowYear())

	// 出力ディレクトリの確認
	if err := os.MkdirAll(outputDir, 0777); err != nil {
		log.Fatal("outputdir error")
		panic(err)
	}

	// jsonデータを出力
	file, err := os.Create(filepath.Join(outputDir, nowMonth+".json"))
	if err != nil {
		log.Fatal("export error")
		panic(err)
	}
	defer file.Close()

	file.Write(([]byte)(out.String()))
}

func main() {
	apiData := scrapeOffDateTime()
	if isEqual := isEqualDate(apiData.Date); !isEqual {
		generateOffDate(apiData)
	}
}
