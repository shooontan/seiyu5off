package util

import (
	"reflect"
	"testing"
	"time"

	"github.com/tkuchiki/faketime"
)

var (
	jsonData = Api{
		Updated: time.Now().Format("2006-01-02 15:04:05"),
		Date:    []string{"2019-05-04", "2019-05-11", "2019-05-18"},
	}

	date1 = []string{"2019-05-04", "2019-05-11", "2019-05-18"}
	date2 = []string{"2019-05-11", "2019-05-04", "2019-05-18"}
	date3 = []string{}
	date4 = []string{"2019-05-04", "2019-05-11"}
	date5 = []string{"2019-05-11", "2019-05-18"}
	date6 = []string{"2019-05-11", "2019-05-18", "2019-05-22"}
)

func TestIsEqualDate(t *testing.T) {
	ft := faketime.NewFaketime(2019, time.May, 13, 0, 0, 0, 0, time.Local)
	ft.Do()

	if IsEqualDate(date1, &jsonData) != true {
		t.Fatal("date1 should be true")
	}

	if IsEqualDate(date2, &jsonData) != true {
		t.Fatal("date2 should be true")
	}

	if IsEqualDate(date3, &jsonData) != false {
		t.Fatal("date3 should be false")
	}

	if IsEqualDate(date4, &jsonData) != false {
		t.Fatal("date4 should be false")
	}

	if IsEqualDate(date5, &jsonData) != true {
		t.Fatal("date5 should be true")
	}

	if IsEqualDate(date6, &jsonData) != false {
		t.Fatal("date6 should be false")
	}

	ft.Undo()
}

func TestMergeDate(t *testing.T) {
	ft := faketime.NewFaketime(2019, time.May, 13, 0, 0, 0, 0, time.Local)
	ft.Do()

	merged1 := MergeDate(date4, &jsonData)
	expected1 := []string{"2019-05-04", "2019-05-11"}
	if !reflect.DeepEqual(merged1, expected1) {
		t.Fatal("merged1 is invalid value")
	}

	merged2 := MergeDate(date6, &jsonData)
	expected2 := []string{"2019-05-04", "2019-05-11", "2019-05-18", "2019-05-22"}
	if !reflect.DeepEqual(merged2, expected2) {
		t.Fatal("merged2 is invalid value")
	}

	ft.Undo()
}
