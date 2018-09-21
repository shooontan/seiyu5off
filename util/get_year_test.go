package util

import (
	"log"
	"testing"
	"time"
)

func TestGetYear(t *testing.T) {
	nowMonth := time.Now().Format("2006")
	nextMonth := time.Now().AddDate(1, 0, 0).Format("2006")

	year := GetYear("01")
	if year != nowMonth && year != nextMonth {
		log.Print(year)
		t.Fatal("faild test")
	}
}
