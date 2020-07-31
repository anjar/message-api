package helpers

import (
	"encoding/json"
	"github.com/kataras/iris/v12/mvc"
)

type Response struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

func isSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

/**
* set response application/json
*/
func ResponseJson(statusCode int, data interface{}) mvc.Response {
	response := Response{
		Success:    isSuccess(statusCode),
	}

	// check if data is error or success, based on statusCode
	if response.Success {
		response.Data = data
	} else {
		response.Error = data
	}

	responseMarshal, _ := json.Marshal(response)

	return mvc.Response{
		ContentType: "application/json",
		Text:        string(responseMarshal),
	}

}
