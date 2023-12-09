package tools

import "strconv"

func StrToInt64(sNum string) int64 {
	num, _ := strconv.Atoi(sNum)
	return int64(num)
}

func IntToStr(num int) string {
	return strconv.Itoa(num)
}

func StrToBool(str string) bool {
	if str == "true" || str == "True" || str == "TRUE" {
		return true
	}
	return false
}
