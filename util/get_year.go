package util

import "time"

func GetYear(month string) string {
	now := time.Now()
	nowMonth := Str2int(now.Format("01"))
	compMonth := Str2int(month)

	// 来月が来年
	if nowMonth == 12 && compMonth == 1 {
		return now.AddDate(1, 0, 0).Format("2006")
	}

	// 先月が去年
	if nowMonth == 1 && compMonth == 12 {
		return now.AddDate(-1, 0, 0).Format("2006")
	}

	// 年の変化なし
	return now.Format("2006")
}
