package util

import "time"

func GetYear(month string) string {
	targetMonth := Str2int(month)
	nowMonth := Str2int(time.Now().Format("01"))

	// 来月も同じ年
	if nowMonth <= targetMonth {
		return time.Now().Format("2006")
	}

	// 来月が次の年
	now := time.Now()
	return now.AddDate(1, 0, 0).Format("2006")
}
