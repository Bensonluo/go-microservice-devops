package common

import "encoding/json"

// SwapTo 结构体赋值 by json tag
func SwapTo(request, category interface{}) (err error) {
	dataByte, err := json.Marshal(request)
	if err != nil {
		return
	}
	return json.Unmarshal(dataByte, category)
}