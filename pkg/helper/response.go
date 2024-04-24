package helper

import "encoding/json"

func ResponseJSON(s interface{}) []byte {
	b, _ := json.Marshal(s)
	return b
}

type HTTPResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
