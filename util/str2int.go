package util

import "strconv"

func Str2int(t string) int {
	i, _ := strconv.Atoi(t)
	return i
}
