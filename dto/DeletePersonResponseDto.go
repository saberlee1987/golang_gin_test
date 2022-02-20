package dto

import (
	"encoding/json"
	"fmt"
)

type DeletePersonResponseDto struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func (deletePersonDto DeletePersonResponseDto) String() string {
	marshal, err := json.Marshal(deletePersonDto)
	if err != nil {
		return fmt.Sprintf("{\"code\":%d,\"text\":\"%s\"}",
			deletePersonDto.Code, deletePersonDto.Text)
	}
	return string(marshal)
}
