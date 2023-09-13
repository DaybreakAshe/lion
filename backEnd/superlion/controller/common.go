//@program: superlion
//@author: yanjl
//@create: 2023-09-13 13:49
package controller

import (
	"fmt"
	"github.com/u2takey/go-utils/json"
	"log"
	"superlion/service"
)

func GetLoginInfoByC(luser any) *service.GoUserInfo {

	luserBean := &service.GoUserInfo{}

	err := json.Unmarshal([]byte(fmt.Sprintln(luser)), luserBean)

	if err != nil {
		log.Printf("get login json format error %s\n:", err.Error())
		return nil
	}
	return luserBean
}
