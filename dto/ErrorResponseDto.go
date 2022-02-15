package dto

import (
	"encoding/json"
	"fmt"
)

type ErrorResponseDto struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func (errorResponse ErrorResponseDto) String() string {
	marshal, err := json.Marshal(errorResponse)
	if err != nil {
		return fmt.Sprintf("{\"code\":%d,\"text\":\"%s\"}",
			errorResponse.Code, errorResponse.Text)
	}
	return string(marshal)
}
