package util

import "strconv"

func Uint64ToStr(num uint64) string {
	return strconv.FormatUint(num, 10) //10进制转换
}
