package util

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Generate(apiData Api) {
	splitDate := strings.Split(apiData.Date[0], "-")
	year := splitDate[0]
	month := splitDate[1]

	p, _ := os.Getwd()

	// json変換
	outputJSON, err := json.Marshal(apiData)
	if err != nil {
		log.Fatal("json encode error")
		panic(err)
	}
	out := new(bytes.Buffer)
	json.Indent(out, outputJSON, "", "    ")

	outputDir := filepath.Join(p, "dist", "api", year)

	// 出力ディレクトリの確認
	if err := os.MkdirAll(outputDir, 0777); err != nil {
		log.Fatal("outputdir error")
		panic(err)
	}

	// jsonデータを出力
	file, err := os.Create(filepath.Join(outputDir, month+".json"))
	if err != nil {
		log.Fatal("export error")
		panic(err)
	}
	defer file.Close()

	file.Write(([]byte)(out.String()))
}
