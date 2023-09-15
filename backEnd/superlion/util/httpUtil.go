//@program: superlion
//@author: yanjl
//@create: 2023-09-15 15:17
package util

import (
	"github.com/u2takey/go-utils/json"
	"io/ioutil"
	"net/http"
)

// ParseResponse 响应体转map
func ParseResponse(response *http.Response) (map[string]interface{}, error) {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result, err
}
