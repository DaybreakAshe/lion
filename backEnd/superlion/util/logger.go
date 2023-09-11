//@program: superlion
//@author: yanjl
//@create: 2023-09-08 17:19
package util

import (
	"fmt"
	"golang.org/x/text/currency"
)

func PrintLog(s string) {
	fmt.Printf("[log.info]_%s_:_%s\n", currency.Date, s)
}

func PrintError(s string) {
	fmt.Println("[error]_%s_:_%s\n", currency.Date, s)
}
