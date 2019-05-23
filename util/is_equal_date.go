package util

import (
	"reflect"
	"sort"
	"time"
)

// IsEqualDate compares Date array to Old Date array
// If same date, return true. Or not, return false.
func IsEqualDate(offDate []string, jsonData *Api) bool {
	if reflect.DeepEqual(offDate, jsonData.Date) {
		return true
	}

	for _, offDateItem := range offDate {
		if !include(jsonData.Date, offDateItem) {
			return false
		}
	}

	for _, dateItem := range jsonData.Date {
		if include(offDate, dateItem) {
			continue
		}

		if !beforeDate(dateItem) {
			return false
		}
	}

	return true
}

// MergeDate merges Date array.
// Date is sorted.
func MergeDate(offDate []string, jsonData *Api) []string {
	for _, dateItem := range jsonData.Date {
		if include(offDate, dateItem) {
			continue
		}

		if beforeDate(dateItem) {
			offDate = append(offDate, dateItem)
		}
	}

	// sort
	sort.SliceStable(offDate, func(i, j int) bool {
		return offDate[i] < offDate[j]
	})

	return offDate
}

func include(arr []string, key string) bool {
	for _, v := range arr {
		if v == key {
			return true
		}
	}
	return false
}

func beforeDate(date string) bool {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t, _ := time.ParseInLocation("2006-01-02", date, loc)
	now := time.Now()
	return t.Before(now)
}
