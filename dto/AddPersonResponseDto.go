package dto

import (
	"encoding/json"
	"fmt"
)

type AddPersonsResponseDto struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func (addPersonResponse AddPersonsResponseDto) String() string {
	marshal, err := json.Marshal(addPersonResponse)
	if err != nil {
		return fmt.Sprintf("{\"code\":%d,\"text\":\"%s\"}",
			addPersonResponse.Code, addPersonResponse.Text)
	}
	return string(marshal)
}
