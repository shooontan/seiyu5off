package util

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

// Generate exports JSON file.
func Generate(apiData Api, jsonpath string, filename string) {
	p, _ := os.Getwd()
	outputDir := filepath.Join(p, "www", "api", "v1", jsonpath)

	// json変換
	outputJSON, err := json.Marshal(apiData)
	if err != nil {
		log.Fatal("json encode error")
		panic(err)
	}
	out := new(bytes.Buffer)
	json.Indent(out, outputJSON, "", "    ")

	// 出力ディレクトリの確認
	if err := os.MkdirAll(outputDir, 0777); err != nil {
		log.Fatal("outputdir error")
		panic(err)
	}

	// jsonデータを出力
	file, err := os.Create(filepath.Join(outputDir, filename))
	if err != nil {
		log.Fatal("export error")
		panic(err)
	}
	defer file.Close()

	file.Write(([]byte)(out.String()))
}
