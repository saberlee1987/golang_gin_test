package dto

import (
	"encoding/json"
	"fmt"
)

type UpdatePersonsResponseDto struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func (updatePersonDto UpdatePersonsResponseDto) String() string {
	marshal, err := json.Marshal(updatePersonDto)
	if err != nil {
		return fmt.Sprintf("{\"code\":%d,\"text\":\"%s\"}",
			updatePersonDto.Code, updatePersonDto.Text)
	}
	return string(marshal)
}
