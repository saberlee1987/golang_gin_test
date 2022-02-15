package dto

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Firstname    string `json:"firstname"`
	LastName     string `json:"lastname"`
	NationalCode string `json:"nationalCode"`
	Age          int    `json:"age"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
}

func (p Person) String() string {
	marshal, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("{\"firstname\":\"%s\",\"lastname\":\"%s\",\"nationalCode\":\"%s\",\"age\":%d,\"email\":\"%s\",\"mobile\":\"%s\"}",
			p.Firstname, p.LastName, p.NationalCode, p.Age, p.Email, p.Mobile)
	}
	return string(marshal)
}
