package util

import "strconv"

func Uint64ToStr(num uint64) string {
	return strconv.FormatUint(num, 10) //10进制转换
}

func StrToUint64(str string) uint64 {
	i, _ := strconv.ParseUint(str, 10, 64)
	return i
}
