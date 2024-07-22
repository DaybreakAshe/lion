//@program: superlion
//@author: yanjl
//@create: 2023-10-09 13:49
package util

import "strconv"

func StrToInt(val string) int64 {
	num, err := strconv.ParseInt(val, 0, 64)

	if err != nil {
		return 0
	}
	return num
}
