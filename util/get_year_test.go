package util

import (
	"testing"
	"time"

	"github.com/tkuchiki/faketime"
)

func TestGetYear(t *testing.T) {
	// 来月が来年
	f1 := faketime.NewFaketime(2018, time.December, 12, 0, 0, 0, 0, time.Local)
	f1.Do()
	year1 := GetYear("12")
	if year1 != "2018" {
		t.Fatal("year1 must be 2018")
	}
	year2 := GetYear("01")
	if year2 != "2019" {
		t.Fatal("year2 must be 2019")
	}
	f1.Undo()

	// 先月が去年
	f2 := faketime.NewFaketime(2019, time.January, 1, 0, 0, 0, 0, time.Local)
	f2.Do()
	year3 := GetYear("12")
	if year3 != "2018" {
		t.Fatal("year3 must be 2018")
	}
	year4 := GetYear("01")
	if year4 != "2019" {
		t.Fatal("year4 must be 2019")
	}
	f2.Undo()

	// 年の変化なし
	f3 := faketime.NewFaketime(2018, time.June, 1, 0, 0, 0, 0, time.Local)
	f3.Do()
	year5 := GetYear("05")
	if year5 != "2018" {
		t.Fatal("year5 must be 2018")
	}
	year6 := GetYear("07")
	if year6 != "2018" {
		t.Fatal("year6 must be 2018")
	}
	f3.Undo()
}
