package util

type Api struct {
	Updated string   `json:"updated"`
	Date    []string `json:"date"`
}

type DateMap map[string][]string

type Scraping struct {
	Updated string  `json:"updated"`
	Date    DateMap `json:"date"`
}
